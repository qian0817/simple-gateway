package general

import (
	"gateway/pipeline"
	"io/ioutil"
	"net/http"
)

type Echo struct {
}

func (e Echo) Init(_ interface{}) {

}

func (e Echo) Handle(w http.ResponseWriter, request *http.Request, _ pipeline.PipelineChain) {
	w.WriteHeader(200)
	body, _ := ioutil.ReadAll(request.Body)
	_, _ = w.Write(body)
}
