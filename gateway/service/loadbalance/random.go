package loadbalance

import (
	"gateway/upstream"
	"math/rand"
	"net/http"
)

type RandomLoadBalance struct{}

func (*RandomLoadBalance) GetNode(nodes []upstream.Node, _ *http.Request) upstream.Node {
	index := rand.Intn(len(nodes))
	return nodes[index]
}
