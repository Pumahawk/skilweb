package views

import (
	"fmt"
	"html/template"
)

func ProjectDetailsLink(id string) string {
	return fmt.Sprintf("/projects/details/%s", id)
}

func ProjectSearchLink() string {
	return "/projects/search"
}

func LinksFuncMap() template.FuncMap {
	return template.FuncMap{
		"link_projectSearch": ProjectSearchLink,
		"link_projectDetails": ProjectDetailsLink,
	}
}
