package main

import (
	"clock-angle-api/controller"
	"log"
	"net/http"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
	if err != nil {
		log.Fatalln("Could not load .env file")
	}

    controller := controller.NewAngleController()
	http.HandleFunc("/api/clock", controller.GetAngle)
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}
}
