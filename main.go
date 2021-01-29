package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

// HealthCheckHandler
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	res, err := json.Marshal(map[string]interface{}{
		"message": "service unavailable",
	})
	fmt.Println(res)
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
