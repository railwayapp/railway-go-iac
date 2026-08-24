# Thin Railway Infrastructure as Code authoring helpers for Go.

Author `.railway/railway.go` using this module, then `railway config plan` /
`railway config apply`. Prefer **one file per project** that owns the whole
environment. Named partials are a last resort for split repos that cannot
share a file.

```go
package main

import "github.com/railwayapp/railway-go-iac/iac"

func Railway() iac.Project {
  db := iac.Postgres("db")
  web := iac.ServiceNamed("web", iac.ServiceConfig{
    "source": iac.Github("org/app"),
    "start":  "./app",
    "env": map[string]any{
      "DATABASE_URL": db.Env("DATABASE_URL"),
    },
  })
  return iac.ProjectNamed("my-app", []any{db, web})
}
```

Config as Code migration: `railway config migrate --lang go`.

Last resort only: declare `const Partial = "api"` in `railway.go` (same role
as `export const partial` in TypeScript). Do not rename a partial after apply.

See https://docs.railway.com/infrastructure-as-code
