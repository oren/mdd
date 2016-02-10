package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/search", search)

	log.Println("server listening")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok")
}

func search(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
