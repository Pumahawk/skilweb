package controllers

import "github.com/pumahawk/skilweb/services"

type SiteMetadataDTO struct {
	Pages []any `json:"pages"`
}

type SearchPagesDTO struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Filters []any `json:"filters"`
}


type TextFilter struct {
	Type string
	Name string
	Label string
}

func NewTextFilter(name, label string) TextFilter {
	return TextFilter{
		Type: services.TextFilterType,
		Name: name,
		Label: label,
	}
}
