package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"stephenbell.dev/ufo-site/pkg/models/postgresql"

	"github.com/golangcollege/sessions"
	_ "github.com/lib/pq"
)

type contextKey string

const contextKeyIsAuthenticated = contextKey("isAuthenticated")

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	sightings     *postgresql.SightingModel
	templateCache map[string]*template.Template
	users         *postgresql.UserModel
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	dsn := flag.String("dsn", "postgresql://stephen:stephen@localhost:5432/stephendb", "PGSQL data source name")
	secret := flag.String("secret", "7dj.12*y4^skqz)ske@3jskv*s+kd1#2", "Secret key")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Open up a connection to the DB
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	// Initialize the template cache
	templateCache, err := newTemplateCache("./ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Initialize a new Session manager
	session := sessions.New([]byte(*secret))

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		sightings:     &postgresql.SightingModel{DB: db},
		templateCache: templateCache,
		users:         &postgresql.UserModel{DB: db},
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
