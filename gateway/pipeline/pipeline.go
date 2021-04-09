package pipeline

import "net/http"

type Pipeline interface {
	Init(data interface{})
	Handle(w http.ResponseWriter, request *http.Request, chain PipelineChain)
}
