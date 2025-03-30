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

type Views = map[string]*template.Template

type ViewHtml struct {
	Name    string
	Base    []string
	Path    string
	FuncMap template.FuncMap
}

func LoadViews(funcMap template.FuncMap) (views Views) {
	vs := []ViewHtml{
		NewPageDHtml(funcMap, "projects-search", "pages/projects-search.html"),
		NewPageDHtml(funcMap, "projects-details", "pages/projects-details.html"),
		NewPageDHtml(funcMap, "hello", "pages/hello.html"),
		NewPageHtml(funcMap, "generic", "pages/generic.html"),
	}

	views = make(Views)
	for _, v := range vs {
		tmpl, err := v.Template()
		if err != nil {
			log.Fatalf("views init: Unable to load page html [Name=%s] [Path=%s]. %v", v.Name, v.Path, err)
		}
		views[v.Name] = tmpl
	}
	return
}

func NewPageDHtml(funcMap template.FuncMap, name, path string) ViewHtml {
	v := NewPageHtml(funcMap, name, path)
	v.Base = []string{"pages/layout.html", "pages/dashboard-layout.html"}
	return v
}

func NewPageHtml(funcMap template.FuncMap, name, path string) ViewHtml {
	return ViewHtml{
		Base:    []string{"pages/layout.html"},
		Name:    name,
		Path:    path,
		FuncMap: funcMap,
	}
}

func (vh *ViewHtml) Template() (*template.Template, error) {
	tmpl, err := template.New("base").Funcs(vh.FuncMap).ParseFS(pages, vh.Base...)
	if err != nil {
		return nil, fmt.Errorf("view html: Unable to load base [Base=%s]. %w", vh.Base, err)
	}

	tmpl, err = tmpl.ParseFS(pages, vh.Path)
	if err != nil {
		return nil, fmt.Errorf("view html: Unable to load path [Path=%s]. %w", vh.Path, err)
	}

	return tmpl, nil
}

func Render(views Views, wr io.Writer, name string, data any) error {
	tpl := views[name]
	if tpl == nil {
		return fmt.Errorf("views render: Not found template %s", name)
	}

	err := tpl.ExecuteTemplate(wr, "layout.html", data)
	if err != nil {
		return fmt.Errorf("views render: Unable view rendering. %w", err)
	}
	return nil
}
