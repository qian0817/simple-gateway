package auth

import (
	"gateway/pipeline"
	"net/http"
)

type BasicAuth struct {
	Username string
	Password string
}

func (b *BasicAuth) Init(data interface{}) {
	m, ok := data.(map[string]interface{})
	if !ok {
		return
	}
	b.Username = m["Username"].(string)
	b.Password = m["Password"].(string)
}

func (b *BasicAuth) Handle(w http.ResponseWriter, request *http.Request, chain pipeline.PipelineChain) {
	username, password, ok := request.BasicAuth()

	if ok && username == b.Username && password == b.Password {
		chain.DoNext(w, request)
		return
	}
	w.WriteHeader(403)
	_, _ = w.Write([]byte("unauthorizated"))
}
