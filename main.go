package main

import (
	"github/jyami/google_oauth_sample/auth"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/login", auth.GoogleLoginHandler).Methods("GET")
	r.HandleFunc("/auth/google/callback", auth.GoogleCallbackHandler).Methods("GET")
	r.HandleFunc("/success", successHandler).Methods("GET")
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}
func successHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("email:%+w", r.URL.Query().Get("email"))
	log.Println("email:%+w", r.URL.Query().Get("name"))

	data := struct {
		Email string
		Name  string
	}{
		Email: r.URL.Query().Get("email"),
		Name:  r.URL.Query().Get("name"),
	}
	log.Println("data:%+w", data)
	templates.ExecuteTemplate(w, "success.html", data)
}
