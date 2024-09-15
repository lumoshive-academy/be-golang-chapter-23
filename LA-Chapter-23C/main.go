// // example add middleware go-chi
// package main

// import (
// 	"net/http"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/chi/v5/middleware"
// )

// func main() {
// 	r := chi.NewRouter()

// 	// Tambahkan middleware
// 	r.Use(middleware.Logger)
// 	r.Use(middleware.Recoverer)

// 	// Tambahkan rute contoh
// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Welcome to the homepage!"))
// 	})

// 	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
// 		panic("This is a test panic!")
// 	})

// 	http.ListenAndServe(":8080", r)
// }

// example not found handler
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", homePage)
	r.Get("/articles/{articleID}", getArticle)
	// Handle not found routes
	r.NotFound(NotFoundHandler)

	http.ListenAndServe(":8080", r)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Blog Home Page!"))
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "articleID")
	w.Write([]byte("Viewing article with ID: " + articleID))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Sorry, the page you are looking for does not exist.", http.StatusNotFound)
}
