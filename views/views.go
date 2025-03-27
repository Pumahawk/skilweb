package views

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"log"
)

//go:embed pages/*
var pages embed.FS
var views map[string]*template.Template

type ViewHtml struct {
	Name string
	Base []string
	Path string
}

func init() {
	vs := []ViewHtml{
		NewPageDHtml("hello", "pages/hello.html"),
		NewPageHtml("404", "pages/404.html"),
	}

	views = make(map[string]*template.Template)
	for _, v := range vs {
		tmpl, err := v.Template()
		if err != nil {
			log.Fatalf("views init: Unable to load page html [Name=%s] [Path=%s]. %v", v.Name, v.Path, err)
		}
		views[v.Name] = tmpl
	}
}

func NewPageDHtml(name, path string) ViewHtml {
	return ViewHtml{
		Base: []string{"pages/layout.html", "pages/dashboard-layout.html"},
		Name: name,
		Path: path,
	}
}

func NewPageHtml(name, path string) ViewHtml {
	return ViewHtml{
		Base: []string{"pages/layout.html"},
		Name: name,
		Path: path,
	}
}

func (vh *ViewHtml) Template() (*template.Template, error) {
	tmpl, err := template.ParseFS(pages, vh.Base...)
	if err != nil {
		return nil, fmt.Errorf("view html: Unable to load base [Base=%s]. %w", vh.Base, err)
	}

	tmpl, err = tmpl.ParseFS(pages, vh.Path)
	if err != nil {
		return nil, fmt.Errorf("view html: Unable to load path [Path=%s]. %w", vh.Path, err)
	}

	return tmpl, nil
}

func Render(wr io.Writer, name string, data any) error {
	tpl := views[name]
	if tpl == nil {
		return fmt.Errorf("views render: Not found template %s", name)
	}

	err := tpl.Execute(wr, data)
	if err != nil {
		return fmt.Errorf("views render: Unable view rendering. %w", err)
	}
	return nil
}
