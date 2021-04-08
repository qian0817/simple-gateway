package loadbalance

import (
	"gateway/upstream"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRandomLoadBalance(t *testing.T) {
	lb := RandomLoadBalance{}
	nodes := []upstream.Node{
		{Port: 9091}, {Port: 9092}, {Port: 9093},
	}
	assert.True(t, lb.GetNode(nodes, &http.Request{}).Port <= 9093)
	assert.True(t, lb.GetNode(nodes, &http.Request{}).Port > 9090)
}
