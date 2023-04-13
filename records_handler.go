package main

import (
	"database/sql"
	"encoding/json"
	"github.com/lib/pq"
	"log"
	"net/http"
)

func recordsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		queryValues := r.URL.Query()
		names, ok := queryValues["name"]
		if !ok || len(names) < 1 {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		var records []Record

		query := "SELECT name, version, release, location, summary FROM packages WHERE name = ANY($1)"
		rows, err := db.Query(query, pq.Array(names))
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var record Record
			err := rows.Scan(&record.Name, &record.Version, &record.Release, &record.Location, &record.Summary)
			if err != nil {
				log.Println(err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			records = append(records, record)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(records)
	}
}
