# Thin Railway Infrastructure as Code authoring helpers for Go.

Author `.railway/railway.go` using this module, then `railway config plan` /
`railway config apply`. Config as Code migration: `railway config migrate --lang go`.

Multi-repo: declare `const Partial = "api"` in `railway.go` (same role as
`export const partial` in TypeScript).

```go
package main

import "github.com/railwayapp/railway-go-iac/iac"

const Partial = "api"

func Railway() iac.Project {
  web := iac.ServiceNamed("web", iac.ServiceConfig{
    "build": "go build -o app .",
    "start": "./app",
  })
  return iac.ProjectNamed("my-app", []any{web})
}
```

See https://docs.railway.com/infrastructure-as-code
