package database

import (
	"userdata-api/internal/logger"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigration(dbURL string) {
	m, err := migrate.New(
		"file://db/migrations",
		dbURL,
	)

	if err != nil {
		logger.Sugar().Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Sugar().Fatalf("❌ Migration failed: %v", err)
	}

	logger.Sugar().Infoln("✅ Database migrations applied successfully.")
}
