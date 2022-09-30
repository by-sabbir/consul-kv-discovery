package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	consulApi "github.com/by-sabbir/consul-kv-discovery/pkg/consul"
)

var (
	CONSUL_ADDR  = "consul:8500"
	SERVICE_ID   = "go-service"
	SERVICE_NAME = "go-service-test"
	SERVICE_PORT = 8000
	SERVICE_HOST = "0.0.0.0"
)

func init() {
	cli, err := consulApi.NewClient(CONSUL_ADDR)
	if err != nil {
		log.Fatalf("can't initiate consul client: %+v\n", err)
	}
	if err := cli.Register(SERVICE_ID); err != nil {
		log.Println("error registering... ", err)
	}
}

func main() {
	if err := Run(); err != nil {
		log.Fatalf("error starting service: %+v\n", err)
	}
}

func Run() error {
	http.HandleFunc("/health", healthcheck)
	srvString := fmt.Sprintf("%s:%d", SERVICE_HOST, SERVICE_PORT)
	if err := http.ListenAndServe(srvString, nil); err != nil {
		return err
	}
	return nil
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok!",
	})
}
