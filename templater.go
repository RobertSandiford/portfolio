package main

import (
	"gweb"
	"html/template"
	"io"
	"github.com/labstack/echo"
)


// Create a struct to manage templates
type Templater struct {
	templates *template.Template
}

// Templater.Render
// Render a template
func (t *Templater) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


// Load Templates
func processTemplates(templater *Templater) {
	templater.templates = gweb.LoadTemplatesRecursively("html", ".html")
}

// Setup templater struct
func setupTemplater(e *echo.Echo) {
	templater := &Templater{}
	templater.templates = gweb.LoadTemplatesRecursively("html", ".html")
	e.Renderer = templater
}

// Reload templates, for faster development
func middlewareReprocessTemplates(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		processTemplates(e.Renderer.(*Templater))
		next(c)
		return nil
	}
}