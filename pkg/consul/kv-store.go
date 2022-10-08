package consul

import (
	"log"

	"github.com/hashicorp/consul/api"
)

type KVClient struct {
	*api.KV
}

func NewKVClient(c ConsulClient) *KVClient {
	return &KVClient{
		c.KV(),
	}
}

func (k *KVClient) PutKV(key, value string) error {
	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err := k.Put(p, nil)
	if err != nil {
		log.Println("error instergin KV: ", err)
		return err
	}

	return nil
}

func (k *KVClient) GetKV(key string) (string, error) {
	p, _, err := k.Get(key, nil)
	if err != nil {
		log.Println("error getting value from key: ", err)
		return "", nil
	}

	return string(p.Value), nil
}
