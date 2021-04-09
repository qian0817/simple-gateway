package general

import (
	"gateway/pipeline"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEchoPipeline_Handle(t *testing.T) {
	auth := Echo{}

	request, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader("test"))
	w := httptest.NewRecorder()
	auth.Handle(w, request, &pipeline.ApplicationPipelineChain{})
	response, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, string(response), "test")
}
