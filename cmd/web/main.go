package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address (port number)")
	flag.Parse()
	f, err1 := os.OpenFile("./tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer f.Close()

	//infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	//errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	errLog := log.New(f, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)

	app := &application{
		errorLog: errLog,
		infoLog:  infoLog,
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
