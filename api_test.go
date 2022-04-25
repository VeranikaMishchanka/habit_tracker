package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/habits", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("hi")
	})

	log.Println("OK")
	http.ListenAndServe(":4000", router)
}
