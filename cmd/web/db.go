package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DEFAULT_DSN = "host=localhost user=web password=password database=snippetbox port=8082 sslmode=disable TimeZone=Europe/Vienna"

func (app *application) connectToDB(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		app.errorLog.Fatal(err)
	}
	app.db = db
}
