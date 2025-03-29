package views

type GenericViewData struct {
	Title string
	Message string
}

func NewGenericViewData(title, message string) GenericViewData {
	return GenericViewData{
		Title: title,
		Message: message,
	}
}

func ServerErrorData(message string) GenericViewData {
	return NewGenericViewData("Server errror", message)
}

func NotFoundData(message string) GenericViewData {
	return NewGenericViewData("Not found", message)
}
