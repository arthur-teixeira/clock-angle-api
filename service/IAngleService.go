package service

import "clock-angle-api/structs"

type IAngleService interface {
    GetAngle(params *structs.Request) (*structs.Response, error)
}
