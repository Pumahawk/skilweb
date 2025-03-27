package views

import (
	"embed"
	"fmt"
	"io"
	"text/template"
)

//go:embed pages/*
var pages embed.FS

type Views struct {
	Template *template.Template
}


func NewViews() (*Views, error) {
	tpl, err := template.New("").ParseFS(pages, "pages/*")
	if err != nil {
		return nil, fmt.Errorf("views: Unable to load pages %w.", err)
	}
	return &Views{tpl}, nil
}

func (views *Views) Render(wr io.Writer, name string, data any) error {
	err := views.Template.ExecuteTemplate(wr, name, data);
	if err != nil {
		return fmt.Errorf("views render: Unable view rendering. %w", err)
	}
	return nil
}
