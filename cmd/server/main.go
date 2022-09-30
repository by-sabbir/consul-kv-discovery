package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	consulApi "github.com/hashicorp/consul/api"
)

var (
	CONSUL_ADDR  = "103.108.147.41:8500"
	SERVICE_ID   = "go-service"
	SERVICE_NAME = "go-service-test"
	SERVICE_PORT = 8000
	SERVICE_HOST = "localhost"
)

func init() {
	// started_at = time.Now()
	consulConfig := &consulApi.Config{
		Address: CONSUL_ADDR,
	}
	client, err := consulApi.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("error creating consul client: %v", err)
	}

	if err := client.Agent().ServiceDeregister(SERVICE_ID); err != nil {
		log.Println("deregistration status: ", err)
	}

	serviceDefinition := &consulApi.AgentServiceRegistration{
		ID:   SERVICE_ID,
		Name: SERVICE_NAME,
		Port: SERVICE_PORT,
		Check: &consulApi.AgentServiceCheck{
			HTTP:     "http://app:8000/health",
			Interval: "15s",
			Timeout:  "30s",
			Notes:    "for medium blog",
		},
		Meta: map[string]string{
			"host": SERVICE_HOST,
			"port": strconv.Itoa(SERVICE_PORT),
		},
	}
	if err := client.Agent().ServiceRegister(serviceDefinition); err != nil {
		log.Fatalf("error registering service: %+v\n", err)
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
