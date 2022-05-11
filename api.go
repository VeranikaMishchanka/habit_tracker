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

	func GetHabits(w http.ResponseWriter, r *http.Request) {
		db := setupDB()
		rows, err := db.Query("SELECT * FROM habits")

		checkErr(err)

		var habits []Habit
	
		for rows.Next() {
			var habit_id int
			var habit_name string
			var habit_subname string
	
			err = rows.Scan(&habit_id, &habit_name, &habit_subname)
			checkErr(err)
	
			movies = append(habits, Habit{HabitID: habit_id, HabitName: habit_name, HabitSubname: habit_subname})
		}
	
		var response = JsonResponse{Type: "success", Data: habits}
	
		json.NewEncoder(w).Encode(response)
	}

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

	  func DeleteHabit(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
	
		habitID := params["habitid"]
	
		var response = JsonResponse{}
	
		if movieID == "" {
			response = JsonResponse{Type: "error", Message: "You are missing habit id"}
		} else {
			db := setupDB()
	
			_, err := db.Exec("DELETE FROM habits where habit_id = $1", habitID )
	
			checkErr(err)
	
			response = JsonResponse{Type: "success", Message: "The habit has been deleted successfully!"}
		}
	
		json.NewEncoder(w).Encode(response)
	}
	
}
