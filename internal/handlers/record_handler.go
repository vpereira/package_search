package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
  "github.com/vpereira/package_search/internal/types"
)

func RecordHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryValues := r.URL.Query()
		names, ok := queryValues["name"]
		if !ok || len(names) != 1 {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		var record types.Record
		err := db.QueryRow("SELECT name, version, release, location, summary FROM packages WHERE name=$1",
			names[0]).Scan(&record.Name, &record.Version, &record.Release, &record.Location, &record.Summary)
		if err != nil {
			if err == sql.ErrNoRows {
				http.NotFound(w, r)
			} else {
				log.Println(err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(record)
	}
}
