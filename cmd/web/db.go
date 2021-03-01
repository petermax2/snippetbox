package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"nirpet.at/snippetbox/pkg/models"
)

const DEFAULT_DSN = "host=localhost user=web password=password database=snippetbox port=8082 sslmode=disable TimeZone=Europe/Vienna"

func openDB(dsn string, errorLog *log.Logger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		errorLog.Fatal(err)
	}
	return db
}

func (app *application) initModels(db *gorm.DB) {
	app.snippets = &models.SnippetModel{DB: db}

	app.infoLog.Println("Starting migration of the DB model")
	app.snippets.Migrate()
}
