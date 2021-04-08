package service

import (
	"encoding/json"
	"gateway/service/loadbalance"
	"gateway/upstream"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceJsonSerializer(t *testing.T) {
	nodes := []upstream.Node{
		{Port: 9091}, {Port: 9092}, {Port: 9093},
	}
	var service = Service{
		HashOn:      HASH,
		LoadBalance: &loadbalance.RandomLoadBalance{},
		Nodes:       nodes,
	}
	var service2 = Service{}
	jsonData, _ := json.Marshal(service)
	_ = json.Unmarshal(jsonData, &service2)
	assert.Equal(t, service2.HashOn, "")
	assert.Equal(t, service2.LoadBalance, &loadbalance.RandomLoadBalance{})
	assert.Equal(t, service2.Nodes, nodes)

}
