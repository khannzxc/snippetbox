package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type applicaton struct {
	logger *slog.Logger
}

func main() {
	// Register the two new handler functions and corresponding route patterns with
	// the servemux, in exactly the same way that we did before.
	addr := flag.String("addr", "localhost:4000", "HTTP network address")

	flag.Parse()
	//infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &applicaton{
		logger: logger,
	}

	//errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET/snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	// srv := &http.Server{
	// 	Addr:     *addr,
	// 	ErrorLog: errorLog,
	// 	Handler:  mux,
	// }

	// infoLog.Printf("Starting server on %s", *addr)
	// err := srv.ListenAndServe()
	// errorLog.Fatal(err)

	logger.Info("starting server", "addr", addr)
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
