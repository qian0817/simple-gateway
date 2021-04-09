package proxy

import (
	"gateway/upstream"
	"net/http"
)

type Proxy interface {
	Handle(w http.ResponseWriter, request *http.Request, node *upstream.Node)
}
