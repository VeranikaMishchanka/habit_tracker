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
	func AddHabit(w http.ResponseWriter, r *http.Request) {
		habit_name := r.FormValue("habitname")
		habit_subname := r.FormValue("habitsubname")

		var response = JsonResponse{}

		if  habit_name == "" {
			response = JsonResponse{Type: "error", Message: "You are missing habit name"}
		} else {
			db := setupDB()
			var lastInsertID int
			err := db.QueryRow("INSERT INTO habits(habitname, habitsubname) VALUES($1, $2) returning id;", habitname, habitsubname).Scan(&lastInsertID)		
			checkErr(err)

			response = JsonResponse{Type: "success", Message: "The habit has been inserted successfully!"}
			}
		
			json.NewEncoder(w).Encode(response)
		
	  }

	
}
