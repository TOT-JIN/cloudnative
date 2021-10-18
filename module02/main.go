package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/healthz", healthHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, 200)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%s %s %d", r.RemoteAddr, r.RequestURI, http.StatusInternalServerError)
	} else {
		log.Printf("%s %s %d", r.RemoteAddr, r.RequestURI, http.StatusOK)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		log.Printf("%s %s %d", r.RemoteAddr, r.RequestURI, http.StatusNotFound)
		return
	}
	for k, v := range r.Header {
		w.Header().Add(k, v[0])
	}
	version := os.Getenv("VERSION")
	if version == "" {
		version = "v0.0.1"
	}
	w.Header().Add("version", version)
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%s %s %d", r.RemoteAddr, r.RequestURI, http.StatusInternalServerError)
	} else {
		log.Printf("%s %s %d", r.RemoteAddr, r.RequestURI, http.StatusOK)
	}
}
