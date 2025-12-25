package main

import (
	"context"
	"fmt"
	"net/http"
)

//Web Development in Go: Middleware pattern
/*
In Go web development, the Middleware pattern is a way to add reusable logic that runs
before and/or after an HTTP request reaches your main handler.

Think of middleware as a layered pipeline that a request passes through.

Request
  ↓
[ Logging Middleware ]
  ↓
[ Auth Middleware ]
  ↓
[ Final Handler ]
  ↓
Response


*/
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleare bofore function")
		fmt.Println("Method:", r.Method)
		fmt.Println("Path:", r.URL.Path)
		next.ServeHTTP(w, r)
		fmt.Println("middleare After function")
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	fmt.Println("Path:", r.URL.Path)
	w.Write([]byte("hello"))
}

// Authentication Middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authrization")
		if token == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

//Chaining Multiple Middlewares
/***
handler := http.HandlerFunc(helloHandler)
http.Handle("/", chain(
    handler,
    loggingMiddleware,
    authMiddleware,
))

*/
func Chain(h http.Handler, middlerwares ...func(http.Handler) http.Handler) http.Handler {
	for i := 0; i < len(middlerwares); i++ {
		h = middlerwares[i](h)
	}
	return h
}

// Middleware with Context (Advanced)
/*
Middleware often adds data to the request context:

Handler can access it:

user:= r.Context().Value("user")
*/
func ContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "admin")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {

	handleFunction := http.HandlerFunc(helloHandler)
	http.Handle("/", LoggingMiddleware(handleFunction))

	http.ListenAndServe(":8080", nil)
}

// Middleware Details

/*
What is Middleware?
A middleware is a function that:
Takes an http.Handler
Returns a new http.Handler
Wraps extra behavior around the original handler

📌 Common uses:
Logging
Authentication & authorization
Request validation
CORS
Rate limiting
Panic recovery


Middleware Structure



*/
