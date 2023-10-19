package main

import (
	"assignment-3/controllers"
	"assignment-3/database"
	"assignment-3/routers"
	"log"
	"time"
)

func main() {
	PORT := ":8080"

	DB, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer DB.Close()

	go startUpdateDataRoutine()

	// Menjalankan server HTTP
	routers.SetupRouter().Run(PORT)

	select {}
}

func startUpdateDataRoutine() {
	for {
		time.Sleep(15 * time.Second)
		controllers.RequestUpdateData()
	}

}
