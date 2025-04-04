package services

import "context"

type PageType = string
const (
	SearchPageType = "search"
)

type SiteMetadata struct {
	Pages []BackendPage
}

type BackendPage struct {
	Type PageType
}

func GetSiteMetadata(ctx context.Context) (*SiteMetadata, error) {
	m := SiteMetadata{
		Pages: []BackendPage{
			{Type: SearchPageType},
		},
	}
	return &m, nil
}
