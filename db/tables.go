package db

import "database/sql"

func InitTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS todo (
		id uuid PRIMARY KEY NOT NULL,
		label VARCHAR(50) UNIQUE NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT FALSE
	)
	`

	_, err := db.Exec(query)
	return err
}
