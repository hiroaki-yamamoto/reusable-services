package server

import (
	html "html/template"
	txt "text/template"
)

// Server represents template rendering server
type Server struct {
	textTemplate map[string]*txt.Template
	htmlTemplate map[string]*html.Template
}

// NewServer constructs a new instance of Server.
func NewServer() *Server {
	return &Server{
		textTemplate: make(map[string]*txt.Template),
		htmlTemplate: make(map[string]*html.Template),
	}
}

// SetTextTemplate sets a text template.
// Note: If there's the same template name in HTMLTemplate,
// the old one (i.e. HTMLTemplate) is removed and new one (i.e. Text Template)
// will be inserted. If there's the same template name in TextTemplate, it will
// be replaced with new one.
func (me *Server) SetTextTemplate(
	tmpName string,
	tmp *txt.Template,
) {
	if _, ok := me.htmlTemplate[tmpName]; ok {
		me.UnsetHTMLTemplate(tmpName)
	}
	me.textTemplate[tmpName] = tmp
	return
}

// UnsetTextTemplate removes the specified template at tmpName
func (me *Server) UnsetTextTemplate(tmpName string) {
	delete(me.textTemplate, tmpName)
}

// SetHTMLTemplate sets a HTML template.
// Note: If there's the same template name in text template,
// the old one (i.e. text template) is removed and new one (i.e. HTML Template)
// will be inserted. If there's the same template name in HTMLTemplate, it will
// be replaced with new one.
func (me *Server) SetHTMLTemplate(
	tmpName string,
	tmp *html.Template,
) {
	if _, ok := me.textTemplate[tmpName]; ok {
		me.UnsetTextTemplate(tmpName)
	}
	me.htmlTemplate[tmpName] = tmp
	return
}

// UnsetHTMLTemplate removes the specified template at tmpName
func (me *Server) UnsetHTMLTemplate(tmpName string) {
	delete(me.htmlTemplate, tmpName)
}
