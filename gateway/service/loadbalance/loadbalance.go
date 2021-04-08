package loadbalance

import (
	"gateway/upstream"
	"net/http"
)

type LoadBalance interface {
	GetNode(nodes []upstream.Node, r *http.Request) upstream.Node
}
