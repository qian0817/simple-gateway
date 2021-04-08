package router

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapper(t *testing.T) {
	routerMapper := NewRoutingMapper()

	routerMapper.AddRouter(&Router{Path: "/"})
	routerMapper.AddRouter(&Router{Path: "/a"})
	routerMapper.AddRouter(&Router{Path: "/a/b"})
	routerMapper.AddRouter(&Router{Path: "/a/*"})
	routerMapper.AddRouter(&Router{Path: "/a/c"})

	assert.Equal(t, routerMapper.GetRouter("/a"), &Router{Path: "/a"})
	assert.Equal(t, routerMapper.GetRouter("/a/b"), &Router{Path: "/a/b"})
	assert.Equal(t, routerMapper.GetRouter("/b"), &Router{Path: "/*"})
	assert.Equal(t, routerMapper.GetRouter("/"), &Router{Path: "/"})
	assert.Nil(t, routerMapper.GetRouter("/b"))
}
