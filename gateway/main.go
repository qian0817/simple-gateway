package main

import (
	"gateway/pipeline"
	"gateway/pipeline/auth"
	"gateway/pipeline/general"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	RegisterSupportPlugin()
	endpoint := os.Getenv("ETCD_ENDPOINT")
	etcdUsername := os.Getenv("ETCD_USERNAME")
	etcdPassword := os.Getenv("ETCD_PASSWORD")
	client, err := clientv3.New(
		clientv3.Config{
			Endpoints: []string{endpoint},
			Username:  etcdUsername,
			Password:  etcdPassword,
		},
	)
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

func RegisterSupportPlugin() {
	pipeline.Register("basic_auth", &auth.BasicAuth{})
	pipeline.Register("echo", &general.Echo{})
}
