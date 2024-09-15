// // exampel basic route go-chi
// package main

// import (
// 	"net/http"

// 	"github.com/go-chi/chi/v5"
// )

// func main() {
// 	r := chi.NewRouter()

// 	r.Get("/items", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Handling GET request"))
// 	})

// 	r.Post("/items", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Handling POST request"))
// 	})

// 	r.Put("/items/{itemID}", func(w http.ResponseWriter, r *http.Request) {
// 		itemID := chi.URLParam(r, "itemID")
// 		w.Write([]byte("Handling PUT request for item ID: " + itemID))
// 	})

// 	r.Delete("/items/{itemID}", func(w http.ResponseWriter, r *http.Request) {
// 		itemID := chi.URLParam(r, "itemID")
// 		w.Write([]byte("Handling DELETE request for item ID: " + itemID))
// 	})

// 	r.Patch("/items/{itemID}", func(w http.ResponseWriter, r *http.Request) {
// 		itemID := chi.URLParam(r, "itemID")
// 		w.Write([]byte("Handling PATCH request for item ID: " + itemID))
// 	})

// 	r.Options("/items", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Handling OPTIONS request"))
// 	})

// 	http.ListenAndServe(":8080", r)
// }

// example router param
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// Route dengan satu parameter
	r.Get("/users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		w.Write([]byte("User ID: " + userID))
	})
	// Route dengan dua parameter
	r.Get("/users/{userID}/posts/{postID}", func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		postID := chi.URLParam(r, "postID")
		w.Write([]byte("User ID: " + userID + ", Post ID: " + postID))
	})

	// Route dengan parameter dan query
	r.Get("/users/{userID}/posts", func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		page := r.URL.Query().Get("page")
		w.Write([]byte("User ID: " + userID + ", Page: " + page))
	})

	http.ListenAndServe(":8080", r)
}
