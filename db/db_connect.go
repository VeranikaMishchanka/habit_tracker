package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Habit struct {
	Habit_id      int
	Habit_name    string
	Habit_subname *string
}

func Init() *gorm.DB {
	dbURL := "postgres://postgres:3110@localhost:5432/tracker"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	return db
}
