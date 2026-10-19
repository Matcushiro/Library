package api

import (
	"fmt"
	"time"
	"log"
	"net/http"
	"encoding/json"
)


type healthResponse struct {
	Status      string `json:"status"`
	ProjectName string `json:"project_name"`
}

type versionResponse struct {
	Version    string `json:"version"`
	ServerTime string `json:"server_time"`
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, successfully connected")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This page does not exist.", http.StatusNotFound)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	response := versionResponse{
		Version:    "1.0.0",
		ServerTime: time.Now().Format(time.RFC3339),
	}
	JSONresponse, err := json.Marshal(response)
	if err != nil {
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(JSONresponse)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := healthResponse{
		Status:      "available",
		ProjectName: "LibraryAPI",
	}
	JSONresponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(JSONresponse)
}

func Init() {
	http.HandleFunc("/healthcheck", healthHandler)
	http.HandleFunc("/api/v1/info", versionHandler)
	http.HandleFunc("/notfound", notFoundHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
