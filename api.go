package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"vk-extractor/tracker/db"

	"github.com/gorilla/mux"
)

const indexHTML = "view/index.html"

func main() {
	gorm := db.Init()
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(indexHTML)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, ""); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	router.HandleFunc("/habits", func(w http.ResponseWriter, r *http.Request) {
		var h db.Habit
		gorm.First(&h)
		fmt.Println(h)
		json.NewEncoder(w).Encode(h)
	}).Methods("GET")

	log.Println("OK")
	http.ListenAndServe(":4000", router)
}
