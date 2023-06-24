package main

import (
	"clock-angle-api/structs"
	"clock-angle-api/validator"
	"encoding/json"
	"log"
	"math"
	"net/http"
)

const (
    HOURS_PER_ROTATION = 12
    MINUTES_PER_ROTATION = 60
    DEGREES_PER_HOUR = 30
    DEGREES_PER_MINUTE = 6
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
        return
	}
    params, err := validator.ValidateParams(r)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        res, _ := json.Marshal(structs.Error{Message: err.Error()})
        w.Write(res)
        return
    }

    hourDegrees := float64(params.Hours * DEGREES_PER_HOUR) + (float64(params.Minutes) * 0.5)
    minuteDegrees := float64(params.Minutes * DEGREES_PER_MINUTE)

    diff := int(math.Floor(math.Abs(hourDegrees - minuteDegrees)))

    res, _ := json.Marshal(structs.Response{Angle: diff})

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func main() {
	http.HandleFunc("/api/v1/hello", hello)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}
}
