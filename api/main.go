package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Get).Methods("GET")
	router.HandleFunc("/html", GetHTML).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Get(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("HELLO WORLD")
}

func GetHTML(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/sample.html")
	t.Execute(w, map[string]string{"number": "12345"})
}
