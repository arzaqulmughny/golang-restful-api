package app

import (
	"database/sql"
	"golang-restful-api/helpers"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root1234@tcp(localhost:3306)/go_restful_api")
	helpers.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}
