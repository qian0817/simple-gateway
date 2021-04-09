package router

import (
	"gateway/pipeline"
	"gateway/service"
)

type Router struct {
	Name        string
	Labels      map[string]string
	Version     string
	Description string
	Publish     bool
	Host        string
	Path        string
	Methods     map[string]bool
	Service     service.Service
	Plugins     []pipeline.Plugin
	pipelines   []pipeline.Pipeline
}

func (r Router) Pipelines() []pipeline.Pipeline {
	if r.pipelines == nil {
		var ans []pipeline.Pipeline
		for i := 0; i < len(r.Plugins); i++ {
			ans = append(ans, r.Plugins[i].Pipelines())
		}
		r.pipelines = ans
	}
	return r.pipelines
}
