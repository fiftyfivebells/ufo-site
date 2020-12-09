package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"stephenbell.dev/ufo-site/pkg/models/postgresql"

	_ "github.com/lib/pq"
)

type application struct {
	errorLog  *log.Logger
	infoLog   *log.Logger
	sightings *postgresql.SightingModel
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	dsn := flag.String("dsn", "postgresql://stephen:stephen@localhost:5432/stephendb", "PGSQL data source name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Open up a connection to the DB
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog:  errorLog,
		infoLog:   infoLog,
		sightings: &postgresql.SightingModel{DB: db},
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
