package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func workersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := Db.Query("SELECT id, name, email, phone, salary  FROM workers")
		if err != nil {
			logger.ERROR("error while reading SQL")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var workers = []Worker{}

		for rows.Next() {
			var worker Worker
			if err := rows.Scan(&worker.ID, &worker.Name, &worker.Email, &worker.Phone, &worker.Salary); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			workers = append(workers, worker)
		}

		workersJson, err := json.Marshal(workers)
		if err != nil {
			logger.ERROR("error while reading SQL")
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(workersJson)
		return

	case http.MethodPost:
		bBody, err := io.ReadAll(r.Body)

		var k Worker
		if err != nil {
			logger.ERROR("error while read body", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(bBody, &k)
		if err != nil {
			logger.ERROR("error while parse body", err)
			w.WriteHeader(http.StatusInternalServerError)
			return

		}

		_, err = Db.Exec("INSERT INTO workers(name, email, phone, salary) VALUES(?, ?, ?, ?)", k.Name, k.Email, k.Phone, k.Salary)
		if err != nil {
			logger.ERROR("error while exec SQL", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		return
	}
}

func workerHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "workers/")
	id, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		data, err := Db.Query("SELECT id, name, email, phone, salary FROM workers WHERE id = ?", id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var worker Worker

		data.Next()
		if err := data.Scan(&worker.ID, &worker.Name, &worker.Email, &worker.Phone, &worker.Salary); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		workerJson, err := json.Marshal(worker)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(workerJson)
		return
	}
}
