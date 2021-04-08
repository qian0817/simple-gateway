package main

import (
	"fmt"
	"gateway/etcd"
	"gateway/router"
	clientv3 "go.etcd.io/etcd/client/v3"
	"io/ioutil"
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
	url := node.CreateUrl(request.RequestURI)
	r, err := http.NewRequest(request.Method, url, request.Body)
	if err != nil {
		panic(err)
	}
	for k, v := range request.Header {
		for _, value := range v {
			r.Header.Add(k, value)
		}
	}
	response, err := handler.webClient.Do(r)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	for k, v := range response.Header {
		for _, value := range v {
			w.Header().Add(k, value)
		}
	}
	w.WriteHeader(response.StatusCode)
	_, err = w.Write(bytes)
	if err != nil {
		panic(err)
	}
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
