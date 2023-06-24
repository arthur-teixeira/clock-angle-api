package service

import (
	"clock-angle-api/repository"
	"clock-angle-api/structs"
	"math"
)

const (
	DEGREES_PER_HOUR     = 30
	DEGREES_PER_MINUTE   = 6
)

type AngleService struct {
    IAngleService,
    repo repository.IAngleRepository
}

func NewAngleService(repo repository.IAngleRepository) *AngleService {
    return &AngleService{
        repo: repo,
    }
}

func (service AngleService) GetAngle(params *structs.Request) (*structs.Response, error) {
    angle, err := service.repo.GetAngle(params)
    if err != nil {
        return nil, err
    }

    if angle.Id != 0 {
        return &structs.Response{
            Angle: angle.Angle,
        }, nil
    }

    diff := service.calculateAngle(params)

    err = service.repo.SaveAngle(params, diff)
    if err != nil {
        return nil, err
    }

	return &structs.Response{
		Angle: diff,
	}, nil
}

func (service AngleService) calculateAngle(params *structs.Request) int {
    hourDegrees := float64((params.Hours % 12)*DEGREES_PER_HOUR) + (float64(params.Minutes) * 0.5)
    minuteDegrees := float64(params.Minutes * DEGREES_PER_MINUTE)

    innerAngle := math.Floor(math.Abs(hourDegrees - minuteDegrees))
    outerAngle := float64(360 - innerAngle)

    diff := int(math.Min(innerAngle, outerAngle))
    return diff
}
