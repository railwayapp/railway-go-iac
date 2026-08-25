package railway_test

import (
	"fmt"

	railway "github.com/railwayapp/railway-go-sdk"
)

func ExampleProjectNamed() {
	web := railway.ServiceNamed("web", railway.ServiceConfig{
		"start": "./app",
	})
	graph := railway.ProjectNamed("demo", []any{web}).Graph()
	fmt.Println(graph["name"])
	// Output: demo
}
