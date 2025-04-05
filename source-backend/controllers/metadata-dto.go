package controllers

type SiteMetadataDTO struct {
	Pages []PagesDTO `json:"pages"`
}

type PagesDTO struct {
	Type string `json:"type"`
}

