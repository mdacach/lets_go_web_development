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

	// Use the http.ListenAndServe() function to start a new web server. We pass
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit.
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
