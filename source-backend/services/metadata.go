package services

import "context"

type PageType = string
const (
	SearchPageType = "search"
)

type SiteMetadata struct {
	Pages []BackendPage
}

type BackendPage interface {
	Type() PageType
}

type SearchPage struct {
	typep PageType
}

func (page *SearchPage) Type() PageType {
	return page.typep
}

func GetSiteMetadata(ctx context.Context) (*SiteMetadata, error) {
	m := SiteMetadata{
		Pages: []BackendPage{
			&SearchPage{
				typep: SearchPageType,
			},
		},
	}
	return &m, nil
}
