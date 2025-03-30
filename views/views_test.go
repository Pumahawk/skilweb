package views

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/pumahawk/skilweb/services"
)

func TestRender(t *testing.T) {
	v := LoadViews(LinksFuncMap())
	data := ProjectsDetailsResponse{
		Title: "Myt project title",
		Data: services.ProjectDetails{
			Name: "MyTest Project",
		},
	}
	var w bytes.Buffer
	err := Render(v, &w, "projects-details", data)
	if err != nil {
		t.Errorf("Invalid render: %v", err)
	}

	ts := w.String()
	if !strings.Contains(ts, "MyTest Project") {
		t.Errorf("Don't cointains project name")
	}
	fmt.Println(ts)
}
