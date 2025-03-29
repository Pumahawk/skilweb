package services

type Project struct {
	Id string
	Name string
}

func ProjectSerach() ([]Project, error) {
	data := []Project{
		{
			Id: "0001",
			Name: "fyrst-project",
		},
	}
	return data, nil
}
