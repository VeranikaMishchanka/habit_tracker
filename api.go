package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"vk-extractor/habit_tracker/db"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

const indexHTML = "view/index.html"

func main() {
	gorm := db.Init()
	router := mux.NewRouter()
	/*router.HandleFunc("/habits", handlers.GetAllHabits).Methods(http.MethodGet)*/
	router.HandleFunc("/habits", handlers.AddHabit).Methods(http.MethodPost)

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

	/*func GetAllHabits(c *gin.Context) {
		var habits []db.Habit
		db.Find(&habits)
		c.JSON(http.StatusOK, gin.H{"data": habits})
	  }
    */
	func AddHabit(c *gin.Context) {
		habit := db.Habit{Habit_name: input.Habit_name, Habit_subname: input.Habit_subname}
		db.DB.Create(&habit)
	  
		c.JSON(http.StatusOK, gin.H{"data": habit})
	  }

	
}
