package loadbalance

import (
	"gateway/upstream"
	"net/http"
	"strings"
)

type HashLoadBalance struct {
	HashOn string
}

const (
	URI         = "uri"
	SERVER_ADDR = "server_addr"
	COOKIE      = "cookie"
	HEADER      = "header"
)

func (h *HashLoadBalance) GetNode(nodes []upstream.Node, r *http.Request) upstream.Node {
	hashcode := 0
	switch h.HashOn {
	case URI:
		hashcode = hash(r.RequestURI)
	case SERVER_ADDR:
		hashcode = hash(r.RemoteAddr)
	case COOKIE:
		hashcode = hash(r.Header.Get("Cookie"))
	default:
		if strings.HasPrefix(h.HashOn, HEADER) {
			index := h.HashOn[len(HEADER)+1:]
			hashcode = hash(r.Header.Get(index))
		} else {
			panic("unsupported hash type")
		}
	}
	return nodes[hashcode%len(nodes)]
}

func hash(s string) int {
	var h = 0
	for _, c := range s {
		h = 31*h + (int(c) & 0xff)
	}
	return h
}
