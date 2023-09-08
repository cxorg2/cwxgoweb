package api

import (
	"net/http"
	"strings"
)

func rootRoute(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if strings.HasPrefix(path, "/api/") {
		apiRoute(w, r)
	}
	if path == "/" {
		rootPage(w, r)
	}
	if path == "/sleep" {
		sleep1s(w, r)
	}
	if path == "/html" {
		htmlPage(w, r)
	}
	if path == "/getIP" {
		getIP(w, r)
	}
	if path == "/getHostName" {
		getHostName(w, r)
	}

	// w.WriteHeader(404)
	// w.Write([]byte("page no found"))
}

func apiRoute(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/api/readfile" {
		readfile(w, r)
	}
	if path == "/api/env" {
		getEnv(w, r)
	}
	if path == "/api/onelog" {
		oneLog(w, r)
	}
}
