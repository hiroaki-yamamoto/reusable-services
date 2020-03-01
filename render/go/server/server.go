package server

import (
	"github.com/hiroaki-yamamoto/reusable-services/render/go/interfaces"
)

// Server represents template rendering server
type Server struct {
	templates map[string]interfaces.ITemplate
}

// New constructs a new instance of Server.
func New() *Server {
	return &Server{
		templates: make(map[string]interfaces.ITemplate),
	}
}

// SetTemplate sets a template.
func (me *Server) SetTemplate(tmpName string, tmp interfaces.ITemplate) {
	me.templates[tmpName] = tmp
	return
}

// UnsetTemplate removes the specified template at tmpName
func (me *Server) UnsetTemplate(tmpName string) {
	delete(me.templates, tmpName)
}
