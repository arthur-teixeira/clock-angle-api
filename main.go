package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
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
    params := r.URL.Query()

    hours := params.Get("hours") 
    hoursInt, err := strconv.Atoi(hours) 
    if err != nil {
        fmt.Fprintf(w, "Error: Hours should be a number")
        return
    }

    hoursInt = hoursInt % HOURS_PER_ROTATION

    minutes := params.Get("minutes")
    minutesInt, err := strconv.Atoi(minutes) 
    if err != nil {
        fmt.Fprintf(w, "Error: Hours should be a number")
        return
    }

    hourDegrees := float64(hoursInt * DEGREES_PER_HOUR) + (float64(minutesInt) * 0.5)
    minuteDegrees := float64(minutesInt * DEGREES_PER_MINUTE)

    innerAngle := math.Abs(hourDegrees - minuteDegrees)
    outerAngle := 360 - innerAngle

    diff := math.Min(innerAngle, outerAngle)

	w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Difference in angle: %.2f\n", diff)
}

func main() {
	http.HandleFunc("/api/v1/hello", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}
}
