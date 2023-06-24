package repository

import "clock-angle-api/structs"

type IAngleRepository interface {
    GetAngle(params *structs.Request) (*structs.Angle, error) 
    SaveAngle(params *structs.Request, angle int) error 
}
