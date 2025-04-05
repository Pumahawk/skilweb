package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pumahawk/skilweb/server"
	"github.com/pumahawk/skilweb/services"
)

func MetadataController(r *http.Request) server.ControllerResponse[any] {
	m, err := services.GetSiteMetadata(r.Context())
	if err != nil {
		log.Printf("controller metadata: Unable to retrieve site metadata. %v", err)
		return server.MessageResponse(500, "Unable to retrieve site metadata")
	}
	result := mapMetadata(m)
	return server.NewResponse(200, result)
}

func mapMetadata(d *services.SiteMetadata) (result SiteMetadataDTO) {
	for _, p := range d.Pages {
		pdto := mapPageDTO(p)
		result.Pages = append(result.Pages, pdto)
	}
	return
}

func mapPageDTO(p services.BackendPage) any {
	switch p := p.(type) {
	case *services.SearchPage: 
		return mapSearchPage(p)
	default:
		panic(fmt.Errorf("metadata-controller mapPageDTO: Unable map filter. Missing type mapping. Add Type mapping to source code. Type=%T", p))
	}
}

func mapSearchPage(p *services.SearchPage) SearchPagesDTO {
	var filters []any
	for _, f := range p.Filters {
		switch f := f.(type) {
		case *services.TextFilter:
			filters = append(filters, mapTextFilter(f))
		default:
			panic(fmt.Errorf("metadata-controller mapSearchPage: Unable map filter. Missing type mapping. Add Type mapping to source code. Type=%T", f))
		}
	}
	return SearchPagesDTO{
		Id: p.Id(),
		Type: p.Type(),
		Filters: filters,
	}
}

func mapTextFilter(f *services.TextFilter) TextFilter {
	return TextFilter{
		Type: f.Type(),
		Name: f.Name(),
		Label: f.Label(),
	}
}
