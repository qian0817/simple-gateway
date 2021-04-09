package pipeline

import (
	"gateway/proxy"
	"gateway/upstream"
	"net/http"
)

type PipelineChain interface {
	DoNext(w http.ResponseWriter, request *http.Request)
}

type ApplicationPipelineChain struct {
	pos       int
	pipelines []Pipeline
	node      *upstream.Node
	proxy     proxy.Proxy
}

func NewApplicationPipelineChain(pipelines []Pipeline, node *upstream.Node) *ApplicationPipelineChain {
	return &ApplicationPipelineChain{
		pos:       -1,
		pipelines: pipelines,
		node:      node,
		proxy:     &proxy.HttpProxy{},
	}
}

func (a *ApplicationPipelineChain) DoNext(w http.ResponseWriter, request *http.Request) {
	a.pos++
	if a.pos < len(a.pipelines) {
		a.pipelines[a.pos].Handle(w, request, a)
	} else {
		if a.proxy != nil {
			a.proxy.Handle(w, request, a.node)
		}
	}
}
