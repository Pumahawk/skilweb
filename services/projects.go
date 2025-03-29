package services

type Project struct {
	Id string
	Name string
	Description string
}

func ProjectSerach() ([]Project, error) {
	data := []Project{
		{
			Id: "001",
			Name: "first-project",
		},
	}
	return data, nil
}
