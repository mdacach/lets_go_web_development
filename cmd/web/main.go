package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"snippetbox/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog     *log.Logger
	infoLog      *log.Logger
	snippetModel *mysql.SnippetModel
}

func main() {
	// Parse the port address from command-line argument
	// We can specify a default value, a description, and the variable will be
	// automatically parsed to a string
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:1234@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse() // Remember to parse the flags, otherwise you will always be using the default value

	// We can create new loggers
	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime|log.LUTC)

	// We can specify flags to choose what the logger will output
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog:     errorLog,
		infoLog:      infoLog,
		snippetModel: &mysql.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if db.Ping() != nil {
		return nil, err
	}
	return db, nil
}
