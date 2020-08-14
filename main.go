package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", healthHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("[main] Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte("{\"status\":\"UP\"}"))
}
