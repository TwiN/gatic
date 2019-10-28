package main

import (
	"encoding/json"
	"github.com/TwinProduction/gatic/core"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", healthHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("[main][main] Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write(structToJsonBytes(&core.HealthStatus{Status: "UP"}))
}

func structToJsonBytes(obj interface{}) []byte {
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Printf("[main][structToJsonBytes] Unable to marshall object to JSON: %s", err.Error())
	}
	return bytes
}
