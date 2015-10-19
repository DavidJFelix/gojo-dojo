package main

import (
	"net/http"
	//Uncomment when you need UUIDs
	//"code.google.com/p/go-uuid/uuid"
)

func SetJSONContentTypeHeader(w http.ResponseWriter, r *http.Request) {
	//FIXME: application/json; charset=utf-8
}

func SetPlatformResponseHeaders(w http.ResponseWriter, r *http.Request) {
	SetRESTResponseHeaders(w, r)
	//FIXME: Transaction-Id and X-Correlation-Id
	// Transaction-Id should be a unique UUID every time a request comes in
	// X-Correlation-Id should be X-Correlation-Id if it is passed in request headers
	// if the request does not have X-Correlation-Id, mirror Transaction-Id
} 

func SetRESTResponseHeaders(w http.ResponseWriter, r *http.Request) {
	SetStandardHeaders(w, r)
	SetJSONContentTypeHeader(w, r)
}

func SetStandardHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Expires", "-1")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
}

