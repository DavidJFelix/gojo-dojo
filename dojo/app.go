package main

import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/codegangsta/cli"
);

func HelloHandler(w http.ResponseWriter, r *http.Request) {
}


func webServer(c *cli.Context) {
	// We're using mux to simplify routing. It's what David knows.
	router := mux.NewRouter()
	
	// Serve the multiplexer
	http.Handle("/", router)
	http.ListenAndServe(":5000", nil)
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
