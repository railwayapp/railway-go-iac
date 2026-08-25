# Thin Railway Infrastructure as Code authoring helpers for Go.

Module: `github.com/railwayapp/railway-go-sdk` (import name `railway`).

Author `.railway/railway.go` using this module, then `railway config plan` /
`railway config apply`. Prefer **one file per project** that owns the whole
environment. Named partials are a last resort for split repos that cannot
share a file.

```go
package main

import "github.com/railwayapp/railway-go-sdk"

func Railway() railway.Project {
  db := railway.Postgres("db")
  web := railway.ServiceNamed("web", railway.ServiceConfig{
    "source": railway.Github("org/app"),
    "start":  "./app",
    "env": map[string]any{
      "DATABASE_URL": db.Env("DATABASE_URL"),
    },
  })
  return railway.ProjectNamed("my-app", []any{db, web})
}
```

Put a `go.mod` next to `.railway/railway.go` (the CLI `go run`s from that
directory):

```
module railway-config

go 1.22

require github.com/railwayapp/railway-go-sdk v0.2.0
```

Config as Code migration: `railway config migrate --lang go`.

Last resort only: declare `const Partial = "api"` in `railway.go` (same role
as `export const partial` in TypeScript). Do not rename a partial after apply.

See https://docs.railway.com/infrastructure-as-code
