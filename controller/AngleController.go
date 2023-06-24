package controller

import (
	"clock-angle-api/db"
	"clock-angle-api/repository"
	"clock-angle-api/service"
	"clock-angle-api/structs"
	"clock-angle-api/validator"
	"encoding/json"
	"net/http"
)

type AngleController struct {}

func NewAngleController() *AngleController {
    return &AngleController{}
}

func (controller AngleController) GetAngle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
        return
	}
    params, err := validator.ValidateParams(r)
    if err != nil {
        controller.badRequest(w, err.Error())
        return
    }

    conn, err := db.GetConn()
    if err != nil {
        controller.internalServerError(w, err.Error())
        return
    }
    defer conn.Close()

    repo := repository.NewAngleRepository(conn)
    service := service.NewAngleService(repo)

    result, err := service.GetAngle(params)
    if err != nil {
        controller.internalServerError(w, err.Error())
        return
    }

    controller.response(w, result)
}

func (controller AngleController) badRequest(w http.ResponseWriter, message string) {
    w.WriteHeader(http.StatusBadRequest)
    res, _ := json.Marshal(structs.Error{Message: message})
    w.Write(res)
}

func (controller AngleController) internalServerError(w http.ResponseWriter, message string) {
    w.WriteHeader(http.StatusInternalServerError)
    res, _ := json.Marshal(structs.Error{Message: message})
    w.Write(res)
}

func (controller AngleController) response(w http.ResponseWriter, result *structs.Response) {
    res, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
