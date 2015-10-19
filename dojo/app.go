package main

import (
	"encoding/json"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/codegangsta/cli"
	"github.com/codegangsta/negroni"
);

/* A simple healthcheck Handler */
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	SetPlatformResponseHeaders(w, r)

	health := map[string]string{
		"status":"ok",
	}
	json.NewEncoder(w).Encode(health)
}

/* Echo the query params out as a JSON object */
func EchoQueryHandler(w http.ResponseWriter, r *http.Request) {
	SetPlatformResponseHeaders(w, r)

	queryMap := r.URL.Query()
	json.NewEncoder(w).Encode(queryMap)
}

func webServer(c *cli.Context) {
	// We're using mux to simplify routing. It's what David knows.
	router := mux.NewRouter()

	// Use the router to route handlers
	router.HandleFunc("/health", HealthHandler).
		Methods("GET")
	router.HandleFunc("/echo", EchoQueryHandler).
		Methods("GET")
	
	// Serve the multiplexer
	loggingServer := negroni.Classic()
	loggingServer.UseHandler(router)
	loggingServer.Run(":5000")
}	


func main() {
	app := cli.NewApp()
	app.Name = "dojo"
	app.Usage = "Run a webserver for the dojo"
	app.Version = "0.0.1"
	app.Author = "David J Felix, and the dojo folks"
	app.Action = webServer
	app.Flags = []cli.Flag {
	}
	app.Run(os.Args)
}
