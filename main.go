package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from snippetbox"))
}
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("allow", "Post")
		// w.WriteHeader(405)
		// w.Write([]byte("method not allowed"))
		http.Error(w, "method not allowed", 405)
		return
	}
	w.Write([]byte("Create a new Snippet"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view/", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Print("starting server on 3000")
	err := http.ListenAndServe("127.0.0.1:3000", mux)
	log.Fatal(err)
}
