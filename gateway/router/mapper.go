package router

import (
	"gateway/router/trietree"
	"strings"
	"sync"
)

type RoutingMapper struct {
	trieTree *trietree.TrieTree
	mutex    *sync.Mutex
}

func NewRoutingMapper() *RoutingMapper {
	return &RoutingMapper{
		trieTree: trietree.NewTrieTree(),
		mutex:    &sync.Mutex{},
	}
}

func (m *RoutingMapper) AddRouter(router *Router) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	subPaths := subPaths(router.Path)
	m.trieTree.Add(subPaths, router)
}

func (m *RoutingMapper) DelRouter(router *Router) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	subPaths := subPaths(router.Path)
	m.trieTree.Del(subPaths)
}

func (m *RoutingMapper) GetRouter(path string) *Router {
	subPaths := subPaths(path)
	router := m.trieTree.Get(subPaths)
	if router == nil {
		return nil
	}
	return router.(*Router)
}

func subPaths(path string) []string {
	if path == "/" {
		return []string{}
	}
	var start = 0
	var end = len(path)
	if len(path) != 0 && path[0] == '/' {
		start = 1
	}
	if len(path) != 0 && path[len(path)-1] == '/' {
		end = len(path) - 1
	}
	subPaths := strings.Split(path[start:end], "/")
	return subPaths
}
