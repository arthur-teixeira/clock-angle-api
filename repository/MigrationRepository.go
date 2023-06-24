package repository

import "database/sql"

type MigrationRepository struct {
    db *sql.DB
}

func NewMigrationRepository(db *sql.DB) *MigrationRepository {
    return &MigrationRepository{
        db: db,
    }
}

func (repo MigrationRepository) Migrate() error {
	sql := "CREATE TABLE IF NOT EXISTS angles (" +
		"	id SERIAL PRIMARY KEY," +
		"	angle INT NOT NULL," +
		"	hour INT NOT NULL," +
		"	minute INT NOT NULL," +
		"	date DATE NOT NULL" +
		")"
    _, err := repo.db.Exec(sql)

    return err
}
