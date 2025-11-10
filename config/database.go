package config

import (
	"database/sql"
	"userdata-api/internal/database"
	"userdata-api/internal/logger"
)

func InitDB() *sql.DB {

	dbURL := "postgres://postgres:1234@localhost:5432/mydb?sslmode=disable"

	database.RunMigration(dbURL)

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {

		logger.Sugar().Fatal(err)
	}
	return conn
}
