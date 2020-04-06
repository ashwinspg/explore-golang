package db

import (
	"strings"

	"github.com/ashwinspg/explore-golang/config"
	"github.com/ashwinspg/explore-golang/utils"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func filePathWithScheme(path string) string {
	const fileURIScheme = "file:///"
	if strings.HasPrefix(path, fileURIScheme) {
		return path
	}
	return fileURIScheme + path
}

//MigrateUp runs upgrade migrations on the database
func MigrateUp() {
	l := utils.LogEntryWithRef()

	m, err := migrate.New(filePathWithScheme(config.MIGRATION_FILES_PATH), config.DATABASE_URL)
	if err != nil {
		l.WithError(err).Fatal("Failed to initialise migration")
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange && err != migrate.ErrLocked {
		l.WithError(err).Fatal("Failed to migrate")
	}
}
