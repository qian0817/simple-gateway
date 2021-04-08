package upstream

import "fmt"

type Node struct {
	Scheme string
	Host   string
	Port   int16
}

func (n Node) CreateUrl(uri string) string {
	scheme := n.Scheme
	host := n.Host
	port := n.Port
	if scheme == "" {
		scheme = "http"
	}
	if host == "" {
		host = "127.0.0.1"
	}
	if port == 0 {
		port = 80
	}
	return fmt.Sprintf("%s://%s:%d%s", scheme, host, port, uri)
}
