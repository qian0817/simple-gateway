package main

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"net/http"
	"time"
)

func main() {
	client, err := clientv3.New(clientv3.Config{Endpoints: []string{"localhost:2379"}})
	if err != nil {
		panic(err)
	}
	server := http.Server{
		Addr:        ":8080",
		Handler:     NewGatewayHandler(client),
		ReadTimeout: 3 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
