package views

import "github.com/pumahawk/skilweb/services"

type ProjectsSearchResponse = DashboardData[[]services.Project]
type ProjectsDetailsResponse = DashboardData[services.ProjectDetails]

type GenericViewData struct {
	Title   string
	Message string
}

func NewGenericViewData(title, message string) GenericViewData {
	return GenericViewData{
		Title:   title,
		Message: message,
	}
}

func ServerErrorData(message string) GenericViewData {
	return NewGenericViewData("Server errror", message)
}

func NotFoundData(message string) GenericViewData {
	return NewGenericViewData("Not found", message)
}

type DashboardData[T any] struct {
	Title string
	Data  T
}
