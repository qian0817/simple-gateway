package auth

import (
	"gateway/pipeline"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicAuth_Init(t *testing.T) {
	data := make(map[string]interface{})
	data["Username"] = "test"
	data["Password"] = "test"
	auth := BasicAuth{}
	auth.Init(data)
	assert.Equal(t, auth.Username, "test")
	assert.Equal(t, auth.Password, "test")
}

func TestBasicAuthFailed(t *testing.T) {
	auth := BasicAuth{Username: "test", Password: "test"}
	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	auth.Handle(w, request, &pipeline.ApplicationPipelineChain{})
	assert.Equal(t, w.Code, 403)
}

func TestBasicAuthSuccess(t *testing.T) {
	auth := BasicAuth{Username: "test", Password: "test"}
	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	request.SetBasicAuth("test", "test")
	auth.Handle(w, request, &pipeline.ApplicationPipelineChain{})
	assert.Equal(t, w.Code, 200)
}
