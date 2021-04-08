package upstream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode_CreateUrl(t *testing.T) {
	assert.Equal(t, Node{}.CreateUrl("/"), "http://127.0.0.1:80/")
	assert.Equal(t, Node{}.CreateUrl("/a"), "http://127.0.0.1:80/a")
	assert.Equal(t, Node{Scheme: "https", Port: 443}.CreateUrl("/"), "https://127.0.0.1:443/")
}
