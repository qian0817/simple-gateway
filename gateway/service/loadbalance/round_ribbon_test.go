package loadbalance

import (
	"gateway/upstream"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRoundRibbonLoadBalance(t *testing.T) {
	lb := RoundRibbonLoadBalance{}
	nodes := []upstream.Node{
		{Port: 9091}, {Port: 9092}, {Port: 9093},
	}
	assert.Equal(t, lb.GetNode(nodes, &http.Request{}).Port, int16(9091))
	assert.Equal(t, lb.GetNode(nodes, &http.Request{}).Port, int16(9092))
	assert.Equal(t, lb.GetNode(nodes, &http.Request{}).Port, int16(9093))
	assert.Equal(t, lb.GetNode(nodes, &http.Request{}).Port, int16(9091))
}
