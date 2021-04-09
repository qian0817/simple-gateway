package main

import (
	"fmt"
	"gateway/etcd"
	"gateway/pipeline"
	"gateway/router"
	clientv3 "go.etcd.io/etcd/client/v3"
	"net/http"
	"strings"
)

type GatewayHandler struct {
	webClient    *http.Client
	routerMapper *router.RoutingMapper
}

func NewGatewayHandler(client *clientv3.Client) *GatewayHandler {
	handler := &GatewayHandler{
		webClient:    &http.Client{},
		routerMapper: etcd.NewEtcdRoutingMapper(client),
	}
	return handler
}

func (handler *GatewayHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	route := handler.routerMapper.GetRouter(getPath(request))
	if route == nil {
		w.WriteHeader(404)
		_, err := fmt.Fprintf(w, "path %s not found", request.RequestURI)
		if err != nil {
			panic(err)
		}
		return
	}
	node := route.Service.LoadBalance.GetNode(route.Service.Nodes, request)
	chain := pipeline.NewApplicationPipelineChain(route.Pipelines(), &node)
	chain.DoNext(w, request)
}

func getPath(request *http.Request) string {
	uri := request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	} else {
		return uri[0:index]
	}
}
