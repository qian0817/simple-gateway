package proxy

import (
	"gateway/upstream"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HttpProxy struct {
}

func (h HttpProxy) Handle(w http.ResponseWriter, request *http.Request, node *upstream.Node) {
	u, err := url.Parse(node.CreateUrl(request.RequestURI))
	if err != nil {
		panic(err)
	}
	request.URL = u
	request.RequestURI = ""

	webClient := &http.Client{}
	response, err := webClient.Do(request)
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
