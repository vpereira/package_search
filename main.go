package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

  "github.com/vpereira/package_search/internal/middlewares"
  "github.com/vpereira/package_search/internal/handlers"

	_ "github.com/lib/pq"
)

func main() {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/record", middlewares.EnforceGet(handlers.RecordHandler(db)))
	http.HandleFunc("/records", middlewares.EnforceGet(handlers.RecordsHandler(db)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
