package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key")

// ===== MODELS =====

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	User  string `json:"user"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Claim2 struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// ===== IN-MEMORY STORE =====

var users = map[string]string{}
var todos = []Todo{}
var mu sync.Mutex
var idCounter = 1

// ===== MIDDLEWARE =====

func Chain(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("REQ:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// attach username to context (simple way)
		r.Header.Set("X-User", claims.Username)

		next.ServeHTTP(w, r)
	})
}

// ===== AUTH HANDLERS =====

func Register(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)

	if u.Username == "" || u.Password == "" {
		http.Error(w, "invalid input", 400)
		return
	}

	users[u.Username] = u.Password
	w.Write([]byte("registered"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)

	pass, ok := users[u.Username]
	if !ok || pass != u.Password {
		http.Error(w, "invalid credentials", 401)
		return
	}

	exp := time.Now().Add(time.Hour)

	claims := &Claims{
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString(jwtKey)

	// exp2 := time.Now().Add(time.Hour)
	// claims2 := &Claim2{
	// 	Username: u.Username,
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		ExpiresAt: jwt.NewNumericDate(exp2),
	// 	},
	// }

	// toen := jwt.NewWithClaims(jwt.SigningMethodHS256, claims2)
	// tostr, _ := toen.SignedString(jwtKey)
	// json.NewEncoder(w).Encode(map[string]string{"token": tostr})

	json.NewEncoder(w).Encode(map[string]string{"token": tokenStr})
}

// ===== TODO HANDLER =====

type TodoHandler struct{}

func (h TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user := r.Header.Get("X-User")

	switch r.Method {

	case http.MethodGet:
		var userTodos []Todo
		for _, t := range todos {
			if t.User == user {
				userTodos = append(userTodos, t)
			}
		}
		json.NewEncoder(w).Encode(userTodos)

	case http.MethodPost:
		var t Todo
		json.NewDecoder(r.Body).Decode(&t)

		mu.Lock()
		t.ID = idCounter
		idCounter++
		t.User = user
		todos = append(todos, t)
		mu.Unlock()

		json.NewEncoder(w).Encode(t)

	case http.MethodPut:
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))

		var updated Todo
		json.NewDecoder(r.Body).Decode(&updated)

		for i, t := range todos {
			if t.ID == id && t.User == user {
				todos[i].Title = updated.Title
				todos[i].Done = updated.Done
			}
		}
		w.Write([]byte("updated"))

	case http.MethodPatch:
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))

		var patch map[string]interface{}
		json.NewDecoder(r.Body).Decode(&patch)

		for i, t := range todos {
			if t.ID == id && t.User == user {
				if v, ok := patch["title"]; ok {
					todos[i].Title = v.(string)
				}
				if v, ok := patch["done"]; ok {
					todos[i].Done = v.(bool)
				}
			}
		}
		w.Write([]byte("patched"))

	case http.MethodDelete:
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))

		for i, t := range todos {
			if t.ID == id && t.User == user {
				todos = append(todos[:i], todos[i+1:]...)
				break
			}
		}
		w.Write([]byte("deleted"))

	default:
		http.Error(w, "method not allowed", 405)
	}
}

// ===== MAIN =====

func main() {

	mux := http.NewServeMux()

	// auth routes
	mux.HandleFunc("/register", Register)
	mux.HandleFunc("/login", Login)

	// protected todo routes
	todoHandler := Chain(TodoHandler{}, Logging, JWTAuth)
	mux.Handle("/todos", todoHandler)

	fmt.Println("Server running at :8080")
	http.ListenAndServe(":8080", mux)
}
