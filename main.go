package main

import (
	"clock-angle-api/controller"
	"log"
	"net/http"
)

func main() {
    controller := controller.NewAngleController()
	http.HandleFunc("/api/clock", controller.GetAngle)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}
}
