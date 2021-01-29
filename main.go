package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

// HealthCheckHandler
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	res, err := json.Marshal(map[string]interface{}{
		"app": os.Getenv("APP_NAME"),
		"message": "service available",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(res)
}

func main() {
	http.HandleFunc("/health", HealthCheckHandler)
	fmt.Println("server running in port 2021")
	if err := http.ListenAndServe(":2021", nil); err != nil {
		log.Fatal(err)
	}
}
