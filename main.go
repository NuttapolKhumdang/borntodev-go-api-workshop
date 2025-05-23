package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Worker struct {
	ID     int
	Name   string
	Email  string
	Phone  string
	Salary int
}

var logger = CreateLogger("Server")

var Db *sql.DB
var WorkerList []Worker

func initDatabase() {
	var err error

	Db, err = sql.Open("sqlite3", "webservice.db")

	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec(`
	CREATE TABLE IF NOT EXISTS workers (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		phone TEXT,
		salary INTEGER
	)`) // create a database table if not exists

	if err != nil {
		logger.ERROR("error while create database table")
		log.Fatal(err)
	}
}

func main() {
	// setup logger
	logger.INFO("initial database")
	initDatabase() // initial database

	defer Db.Close()
	defer logger.INFO("close database connection")

	// initial for route `/workers`
	http.Handle("/workers", useCorsMiddleware(http.HandlerFunc(workersHandler)))
	// initial for route `/workers/:id`
	http.Handle("/workers/", useCorsMiddleware(http.HandlerFunc(workerHandler)))

	logger.INFO("listen for ::1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
