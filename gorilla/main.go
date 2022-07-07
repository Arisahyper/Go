package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()	// *Router型の変数r

	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", logoutHandler)

	http.Handle("/", r)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

var homeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
})

var loginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
})

var logoutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
})
