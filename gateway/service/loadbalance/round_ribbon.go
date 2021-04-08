package loadbalance

import (
	"gateway/upstream"
	"net/http"
)

type RoundRibbonLoadBalance struct {
	index int
}

func (rr *RoundRibbonLoadBalance) GetNode(nodes []upstream.Node, _ *http.Request) upstream.Node {
	node := nodes[rr.index]
	rr.index = (rr.index + 1) % len(nodes)
	return node
}
