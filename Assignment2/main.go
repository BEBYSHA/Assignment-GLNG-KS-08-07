package main

import (
	"assignment2/data"
	"assignment2/routers"
)

func main() {
	PORT := ":8080"
	_, err := data.InitDB()

	if err != nil {
		panic(err)
	}

	// migrate database

	data.Migrate()
	routers.SetupRouter().Run(PORT)
}
