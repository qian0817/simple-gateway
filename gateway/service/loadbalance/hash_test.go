package loadbalance

import (
	"gateway/upstream"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUriHash(t *testing.T) {
	lb := HashLoadBalance{HashOn: URI}
	nodes := []upstream.Node{
		{Port: 9091}, {Port: 9092}, {Port: 9093},
	}
	assert.True(t, lb.GetNode(nodes, &http.Request{}).Port <= 9093)
	assert.True(t, lb.GetNode(nodes, &http.Request{}).Port > 9090)
	// same hash
	assert.Equal(t, lb.GetNode(nodes, &http.Request{}), lb.GetNode(nodes, &http.Request{}))
	// different hash
	assert.NotEqual(
		t,
		lb.GetNode(nodes, &http.Request{RequestURI: "/a"}),
		lb.GetNode(nodes, &http.Request{RequestURI: "/b"}),
	)
}
