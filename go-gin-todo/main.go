package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
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

// ===== IN-MEMORY STORE =====

var users = map[string]string{}
var todos = []Todo{}
var mu sync.Mutex
var idCounter = 1

// ===== MIDDLEWARE =====

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		tokenStr := auth[len("Bearer "):]

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("user", claims.Username)
		c.Next()
	}
}

// ===== AUTH HANDLERS =====

func Register(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil || u.Username == "" || u.Password == "" {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	users[u.Username] = u.Password
	c.JSON(200, gin.H{"message": "registered"})
}

func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	pass, ok := users[u.Username]
	if !ok || pass != u.Password {
		c.JSON(401, gin.H{"error": "invalid credentials"})
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

	c.JSON(200, gin.H{"token": tokenStr})
}

// ===== TODO HANDLERS =====

func GetTodos(c *gin.Context) {
	user := c.GetString("user")

	var userTodos []Todo
	for _, t := range todos {
		if t.User == user {
			userTodos = append(userTodos, t)
		}
	}

	c.JSON(200, userTodos)
}

func CreateTodo(c *gin.Context) {
	user := c.GetString("user")

	var t Todo
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(400, gin.H{"error": "invalid body"})
		return
	}

	mu.Lock()
	t.ID = idCounter
	idCounter++
	t.User = user
	todos = append(todos, t)
	mu.Unlock()

	c.JSON(201, t)
}

func GetTodoByID(c *gin.Context) {
	user := c.GetString("user")
	id, _ := strconv.Atoi(c.Param("id"))

	for _, t := range todos {
		if t.ID == id && t.User == user {
			c.JSON(200, t)
			return
		}
	}

	c.JSON(404, gin.H{"error": "not found"})
}

func UpdateTodo(c *gin.Context) {
	user := c.GetString("user")
	id, _ := strconv.Atoi(c.Param("id"))

	var updated Todo
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(400, gin.H{"error": "invalid body"})
		return
	}

	for i, t := range todos {
		if t.ID == id && t.User == user {
			todos[i].Title = updated.Title
			todos[i].Done = updated.Done
			c.JSON(200, gin.H{"message": "updated"})
			return
		}
	}

	c.JSON(404, gin.H{"error": "not found"})
}

func PatchTodo(c *gin.Context) {
	user := c.GetString("user")
	id, _ := strconv.Atoi(c.Param("id"))

	var patch map[string]interface{}
	if err := c.ShouldBindJSON(&patch); err != nil {
		c.JSON(400, gin.H{"error": "invalid body"})
		return
	}

	for i, t := range todos {
		if t.ID == id && t.User == user {

			if v, ok := patch["title"]; ok {
				if title, ok := v.(string); ok {
					todos[i].Title = title
				}
			}

			if v, ok := patch["done"]; ok {
				if done, ok := v.(bool); ok {
					todos[i].Done = done
				}
			}

			c.JSON(200, gin.H{"message": "patched"})
			return
		}
	}

	c.JSON(404, gin.H{"error": "not found"})
}

func DeleteTodo(c *gin.Context) {
	user := c.GetString("user")
	id, _ := strconv.Atoi(c.Param("id"))

	for i, t := range todos {
		if t.ID == id && t.User == user {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(200, gin.H{"message": "deleted"})
			return
		}
	}

	c.JSON(404, gin.H{"error": "not found"})
}

// ===== MAIN =====

func main() {

	r := gin.Default()

	// public routes
	r.POST("/register", Register)
	r.POST("/login", Login)

	// protected routes
	auth := r.Group("/todos")
	auth.Use(JWTAuthMiddleware())
	{
		auth.GET("", GetTodos)
		auth.POST("", CreateTodo)
		auth.GET("/:id", GetTodoByID)
		auth.PUT("/:id", UpdateTodo)
		auth.PATCH("/:id", PatchTodo)
		auth.DELETE("/:id", DeleteTodo)
	}

	r.Run(":8080")
}
