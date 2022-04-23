package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Parse the port address from command-line argument
	// We can specify a default value, a description, and the variable will be
	// automatically parsed to a string
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse() // Remember to parse the flags, otherwise you will always be using the default value

	// We can create new loggers
	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime|log.LUTC)

	// We can specify flags to choose what the logger will output
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
