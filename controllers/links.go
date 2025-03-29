package controllers

import (
	"fmt"
	"html/template"
)

func ProjectDetailsLink(id string) string {
	return fmt.Sprintf("/projects/details/%s", id)
}

func LinksFuncMap() template.FuncMap {
	return template.FuncMap{
		"link_projectDetails": ProjectDetailsLink,
	}
}
