package database

import (
	"assignment-3/helpers"
	"log"
)

// percobaan saja
func InsertData() {
	if err := DB.QueryRow(`INSERT INTO weather(water, wind) 
	VALUES($1, $2)`, helpers.RandomNumber(1, 100), helpers.RandomNumber(1, 100)); err != nil {
		log.Fatal(err)
	}
}
