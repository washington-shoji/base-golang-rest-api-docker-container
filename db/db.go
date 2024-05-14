package db

import (
	"base-golang-rest-api-docker-container/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDBConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.EnvConfig("DB_HOST"), config.EnvConfig("DB_USER"), config.EnvConfig("DB_PASSWORD"), config.EnvConfig("DB_NAME"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
