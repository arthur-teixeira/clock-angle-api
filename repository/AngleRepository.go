package repository

import (
	"clock-angle-api/structs"
	"database/sql"
	"time"
)

type AngleRepository struct {
    IAngleRepository,
    db *sql.DB
}

func NewAngleRepository(db *sql.DB) *AngleRepository {
    return &AngleRepository{
        db: db,
    }
}


func (repo AngleRepository) GetAngle(params *structs.Request) (*structs.Angle, error) {
    var angle structs.Angle
    row := repo.db.QueryRow("SELECT id, angle FROM angles WHERE hour = $1 AND minute = $2", params.Hours, params.Minutes)
    if err := row.Err(); err != nil {
        return nil, err
    }
    row.Scan(&angle.Id, &angle.Angle)

    return &angle, nil
}

func (repo AngleRepository) SaveAngle(params *structs.Request, angle int) error {
    now := time.Now().Format("2006-01-02")
    _, err := repo.db.Exec("INSERT INTO angles (hour, minute, angle, date) VALUES ($1, $2, $3, $4)", params.Hours, params.Minutes, angle, now)
    if err != nil {
        return err
    }

    return nil
}
