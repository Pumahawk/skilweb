package views

import (
	"embed"
	"fmt"
	"html/template"
	"io"
)

//go:embed pages/*
var pages embed.FS
var views []View = []View{
	{
		Name: "hello",
		Path: []string{"pages/layout.html", "pages/hello.html"},
	},
	{
		Name: "404",
		Path: []string{"pages/layout.html", "pages/404.html"},
	},
}

type Templates = map[string]*template.Template
type Views struct {
	Templates Templates
}

type View struct {
	Name string
	Path []string
}

func NewViews() (*Views, error) {
	tpls := make(Templates)
	for _, vi := range views {
		tpl, err := template.ParseFS(pages, vi.Path...)
		if err != nil {
			return nil, fmt.Errorf("views: Unable to load pages %w.", err)
		}
		tpls[vi.Name] = tpl
	}
	return &Views{tpls}, nil
}

func (views *Views) Render(wr io.Writer, name string, data any) error {
	tpl := views.Templates[name]
	if tpl == nil {
		return fmt.Errorf("views render: Not found template %s", name)
	}
	err := tpl.Execute(wr, data)
	if err != nil {
		return fmt.Errorf("views render: Unable view rendering. %w", err)
	}
	return nil
}
