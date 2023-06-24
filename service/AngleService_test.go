package service

import (
	"clock-angle-api/repository"
	"clock-angle-api/structs"
	"testing"
)

type MockAngleRepository struct {
    repository.IAngleRepository
    angle structs.Angle
    inserted bool
}

func newMockAngleRepository() *MockAngleRepository {
    return &MockAngleRepository{}
}

func (repo MockAngleRepository) SetAngle(angle int) {
    repo.angle.Id = 1
    repo.angle.Angle = angle
}

func (repo MockAngleRepository) GetAngle(params *structs.Request) (*structs.Angle, error) {
    return &repo.angle, nil
}

func (repo *MockAngleRepository) SaveAngle(params *structs.Request, angle int) error {
    repo.inserted = true
    return nil
}

func TestAngleInDb(t *testing.T) {
    repo := newMockAngleRepository()
    repo.SetAngle(90)
    service := NewAngleService(repo)
    params := &structs.Request{
        Hours: 3,
        Minutes: 0,
    }


    angle, _ := service.GetAngle(params)
    if angle.Angle != 90 {
        t.Errorf("Expected angle to be 90, got %v", angle.Angle)
    }

    if (repo.inserted) {
        t.Errorf("Expected angle to be fetched from db, not inserted")
    }
}

func TestAngleNotInDb(t *testing.T) {
    repo := newMockAngleRepository()
    service := NewAngleService(repo)
    params := &structs.Request{
        Hours: 3,
        Minutes: 0,
    }

    angle, _ := service.GetAngle(params)
    if angle.Angle != 90 {
        t.Errorf("Expected angle to be 90, got %v", angle.Angle)
    }

    if (!repo.inserted) {
        t.Errorf("Expected angle to be inserted into db")
    }
}
