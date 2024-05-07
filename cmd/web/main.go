package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"tournamentsupport.com/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	repo     *models.Repository
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address (port number)")
	dsn := flag.String("dsn", "file:main.db", "SQLite data source name")
	flag.Parse()
	f, err1 := os.OpenFile("./tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer f.Close()

	// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	errLog := log.New(f, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)

	db, errDB := openDB(*dsn)

	if errDB != nil {
		errLog.Fatal(errDB)
	}
	defer db.Close()

	app := &application{
		errorLog: errLog,
		infoLog:  infoLog,
		repo:     &models.Repository{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s \n", *addr)
	err := srv.ListenAndServe()
	errLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", dsn)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil

}
