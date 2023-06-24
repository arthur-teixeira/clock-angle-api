package db

import (
	"database/sql"
	"fmt"
    _ "github.com/lib/pq"
)

const (
    host = "db"
    port = 5432
    user = "postgres"
    password = "root"
    dbname = "angleapi"
)

func GetConn() (*sql.DB, error) {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return db, nil
}
