package main

import (
	"bufio"
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

/* Simply go get the github activity feed for a user and return it */
func ActivityHandler(w http.ResponseWriter, r *http.Request) {
	SetPlatformResponseHeaders(w, r)
	
	// Get request variables and format upstream URL
	vars :=	mux.Vars(r)
	URL := "https://api.github.com/users/" + vars["githubUsername"] + "/events"

	response, err := http.Get(URL)
	if err == nil {
		bufio.NewReader(response.Body).WriteTo(w)
	}
}

type FanOutResponse struct {
	url string
	response *http.Response
	err error
}

/* Demonstrate fan-out by calling httpbin 12 times all at once */
func HTTPBinHandler(w http.ResponseWriter, r *http.Request) {
	SetPlatformResponseHeaders(w, r)

	// Create a channel and response holder
	ch := make(chan *FanOutResponse)
	responses := []*FanOutResponse{}

	// These are the status codes we'll hit
	statusCodes := [12]string{
		"200", "201", "204", "206",
		"301", "307",
		"400", "401", "403", "404",
		"500", "501",
	}

	baseURL := "https://httpbin.org/status/"

	// Place the requests
	for _, code := range statusCodes {
		// Create goroutine for the url
		go func(url string) {
			resp, err := http.Get(url)
			// Write our response to the channel
			ch <- &FanOutResponse{url, resp, err}
		}(baseURL + code)
	}

	// Await the response object, this could use optimization
	waitLoop:
		for {
			r := <-ch
			responses = append(responses, r)
			if len(responses) == len(statusCodes) {
				break waitLoop
			}
		}

	//FIXME: Replace this with the json responses
	done := map[string]string{
		"done?":"yep",
	}
	json.NewEncoder(w).Encode(done)
}

// FIXME: create this endpoint
/* This endpoint compares multiple user's github activity streams
and declares a winner based on who has the most "PushEvents"
*/
func PushActivityFightHandler(w http.ResponseWriter, r *http.Request) {
}
	

func webServer(c *cli.Context) {
	// We're using mux to simplify routing. It's what David knows.
	router := mux.NewRouter()

	// Use the router to route handlers
	router.HandleFunc("/health", HealthHandler).
		Methods("GET")
	router.HandleFunc("/echo", EchoQueryHandler).
		Methods("GET")
	router.HandleFunc("/activity/{githubUsername}", ActivityHandler).
		Methods("GET")
	router.HandleFunc("/httpbin", HTTPBinHandler).
		Methods("Get")
	
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
