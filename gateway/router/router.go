package router

import (
	"gateway/plugin"
	"gateway/service"
)

type Router struct {
	Name        string
	Labels      map[string]string
	Version     string
	Description string
	Publish     bool
	Host        string
	Path        string
	Methods     map[string]bool
	Service     service.Service
	Plugins     []plugin.Plugin
}
