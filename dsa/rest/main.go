/**

Request
// {
//   "users": [
//     {"name": "Vicky", "age": 25},
//     {"name": "", "age": 30},
//     {"name": "John", "age": -5}
//   ]
// }


Response
// {
//   "valid_users": [
//     {"name": "Vicky", "age": 25}
//   ]
// }

*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Users []User `json:"valid_users"`
}
type Request struct {
	Users []User `json:"users"`
}

type UserHandler struct{}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware called")

		fmt.Println("Method : ", r.Method)
		fmt.Println("URL : ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func Chain(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range m {
		h = middleware(h)
	}
	return h
}

func (h UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method Not Allowed ", http.StatusMethodNotAllowed)
		return
	}
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	var ValidUsers []User

	for _, user := range req.Users {

		if user.Name != "" && user.Age > 0 {
			ValidUsers = append(ValidUsers, user)
		}
	}

	res := Response{Users: ValidUsers}
	json.NewEncoder(w).Encode(res)
}

func main() {
	mux := http.NewServeMux()
	userhandler := UserHandler{}
	mux.Handle("/users", LoggingMiddleware(userhandler))
	fmt.Println("Server running at 8080")
	http.ListenAndServe(":8080", mux)
}
