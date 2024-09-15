// // method not allowed handler
// package main

// import (
// 	"net/http"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/chi/v5/middleware"
// )

// func main() {
// 	r := chi.NewRouter()
// 	r.Use(middleware.Logger)

// 	r.Get("/users", getUsers)               // Mendukung GET
// 	r.Post("/users", createUser)            // Mendukung POST
// 	r.Put("/users/{userID}", updateUser)    // Mendukung PUT
// 	r.Delete("/users/{userID}", deleteUser) // Mendukung DELETE

// 	// Handle not allowed methods
// 	r.MethodNotAllowed(MethodNotAllowedHandler)

// 	http.ListenAndServe(":8080", r)
// }

// func getUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Get all users"))
// }

// func createUser(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Create a new user"))
// }

// func updateUser(w http.ResponseWriter, r *http.Request) {
// 	userID := chi.URLParam(r, "userID")
// 	w.Write([]byte("Update user with ID: " + userID))
// }

// func deleteUser(w http.ResponseWriter, r *http.Request) {
// 	userID := chi.URLParam(r, "userID")
// 	w.Write([]byte("Delete user with ID: " + userID))
// }

// func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// }

// example middleware customer
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(CustomHeaderMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	http.ListenAndServe(":8080", r)
}

func CustomHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Custom-Header") == "" {
			http.Error(w, "X-Custom-Header is required", http.StatusBadRequest)
			return
		}
		log.Println("X-Custom-Header is present")
		next.ServeHTTP(w, r)
	})
}
