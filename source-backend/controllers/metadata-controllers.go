package controllers

import (
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
		result.Pages = append(result.Pages, PagesDTO{p.Type})
	}
	return
}
