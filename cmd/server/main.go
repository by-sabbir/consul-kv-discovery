package main

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/by-sabbir/consul-kv-discovery/internal/consul"
)

func main() {
	if err := Run(); err != nil {
		log.Fatalf("error starting service: %+v\n", err)
	}
}

func Run() error {
	http.HandleFunc("/health", healthcheck)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		return err
	}
	return nil
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok!",
	})
}
