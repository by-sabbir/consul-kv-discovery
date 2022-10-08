package consul

import (
	"fmt"
	"log"

	consulApi "github.com/by-sabbir/consul-kv-discovery/pkg/consul"
)

type ServiceDefinition struct {
	ConsulAddr  string
	ServiceId   string
	ServiceHost string
	ServicePort int
}

func init() {
	var sd ServiceDefinition
	sd.ConsulAddr = "consul:8500"
	sd.ServiceId = "go-service"
	sd.ServiceHost = "0.0.0.0"
	sd.ServicePort = 8000

	cli, err := consulApi.NewClient(sd.ConsulAddr)
	if err != nil {
		log.Fatalf("can't initiate consul client: %+v\n", err)
	}
	if err := cli.Register(sd.ServiceId); err != nil {
		log.Println("error registering... ", err)
	}
	log.Println("service registered: ", sd.ServiceId)

	consulStore := consulApi.NewKVClient(cli)

	srvString := fmt.Sprintf("%s:%d", sd.ServiceHost, sd.ServicePort)

	if err := consulStore.PutKV("apigw", "apigw.example.com"); err != nil {
		log.Println("error creating key: ", err)

	}

	if err := consulStore.PutKV(sd.ServiceId, srvString); err != nil {
		log.Println("error creating key: ", err)

	}

	url, err := consulStore.GetKV("apigw")
	if err != nil {
		log.Println("could not get kv for apigw: ", err)
	}
	log.Println("apigw baseUrl: ", url)

}
