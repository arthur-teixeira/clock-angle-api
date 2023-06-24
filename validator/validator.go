package validator

import (
	"clock-angle-api/structs"
	"errors"
	"net/http"
	"strconv"
)

func ValidateParams(r *http.Request) (*structs.Request, error) {
    params := r.URL.Query()

    hours, err := validateHours(params.Get("hours"))
    if err != nil {
        return nil, err
    }

    minutes, err := validateMinutes(params.Get("minutes"))
    if err != nil {
        return nil, err
    }

    return &structs.Request{
        Hours: hours,
        Minutes: minutes,
    }, nil
}

func validateMinutes(minutes string) (int, error) {
    if minutes == "" {
        return 0, nil
    }

    minutesInt, err := strconv.Atoi(minutes) 
    if err != nil {
        return 0, errors.New("Minutes should be a number")
    }

    if minutesInt > 59 {
        return 0, errors.New("The maximum value for minutes is 60")
    }

    return minutesInt, nil
}

func validateHours(hours string) (int, error) {
    if hours == "" {
        return 0, errors.New("Hours is required")
    }

    hoursInt, err := strconv.Atoi(hours) 
    if err != nil {
        return 0, errors.New("Hours should be a number")
    }

    if hoursInt > 24 {
        return 0, errors.New("The maximum value for hours is 24")
    }

    return hoursInt, nil
}
