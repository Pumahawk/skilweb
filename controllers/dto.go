package controllers

import "github.com/pumahawk/skilweb/services"

type ProjectsSearchResponse = DashboardData[[]services.Project]

type DashboardData[T any] struct {
	Title string
	Data T
}
