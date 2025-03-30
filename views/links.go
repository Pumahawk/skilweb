package views

import (
	"fmt"
	"html/template"
)

func ProjectDetailsLink(id string) string {
	return fmt.Sprintf("/projects/%s", id)
}

func ProjectCreateLink() string {
	return "/projects"
}

func ProjectSearchLink() string {
	return "/projects/search"
}

func LinksFuncMap() template.FuncMap {
	return template.FuncMap{
		"link_projectSearch": ProjectSearchLink,
		"link_projectDetails": ProjectDetailsLink,
		"link_projectCreate": ProjectCreateLink,
	}
}
