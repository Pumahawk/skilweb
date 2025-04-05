package services

import "context"

type PageType = string

const (
	SearchPageType = "search"
)

const (
	TextFilterType = "text"
)

type Filter interface {
	Name() string
	Label() string
	Type() string
}

type TextFilter struct {
	Namef  string
	Labelf string
}

func (f TextFilter) Name() string {
	return f.Namef
}

func (f TextFilter) Label() string {
	return f.Labelf
}

func (TextFilter) Type() string {
	return TextFilterType
}

type SiteMetadata struct {
	Pages []BackendPage
}

type BackendPage interface {
	Id() string
	Type() PageType
}

type SearchPage struct {
	Idf string
	Typef   PageType
	Filters []Filter
}

func (page *SearchPage) Id() string {
	return page.Idf
}

func (page *SearchPage) Type() PageType {
	return page.Typef
}

func GetSiteMetadata(ctx context.Context) (*SiteMetadata, error) {
	m := SiteMetadata{
		Pages: []BackendPage{
			&SearchPage{
				Idf: "search.projects",
				Typef: SearchPageType,
				Filters: []Filter{
					&TextFilter{
						Namef: "projectName",
						Labelf: "search.projects.projectName",
					},
				},
			},
		},
	}
	return &m, nil
}
