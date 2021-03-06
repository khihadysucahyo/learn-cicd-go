package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// HealthCheckHandler func
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	res, err := json.Marshal(map[string]interface{}{
		"message": "elastic container service available",
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
	godotenv.Load()
	http.HandleFunc("/health", HealthCheckHandler)
	fmt.Printf("%s server running in port 2021", os.Getenv("APP_NAME"))
	if err := http.ListenAndServe(":2021", nil); err != nil {
		log.Fatal(err)
	}
}
