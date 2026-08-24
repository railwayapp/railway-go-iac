package iac_test

import (
	"fmt"

	"github.com/railwayapp/railway-go-iac/iac"
)

func ExampleProjectNamed() {
	web := iac.ServiceNamed("web", iac.ServiceConfig{
		"start": "./app",
	})
	graph := iac.ProjectNamed("demo", []any{web}).Graph()
	fmt.Println(graph["name"])
	// Output: demo
}
