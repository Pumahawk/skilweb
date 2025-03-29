package controllers

import "github.com/pumahawk/skilweb/services"

type ProjectsSearchResponse = DashboardData[[]services.Project]
type ProjectsDetailsResponse = DashboardData[services.ProjectDetails]

type DashboardData[T any] struct {
	Title string
	Data T
}
