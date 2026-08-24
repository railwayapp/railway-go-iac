# Thin Railway Infrastructure as Code authoring helpers for Go.

Author `.railway/railway.go` using this module, then `railway config plan` /
`railway config apply`. Prefer **one file per project** that owns the whole
environment. Named partials are a last resort for split repos that cannot
share a file.

## Install

This is a Go module. There is no package registry: a git tag on this repo
**is** the release.

```bash
go get github.com/railwayapp/railway-go-iac@v0.2.0
```

The CLI evaluates `.railway/railway.go` with `go run`, so the file's module
must be able to resolve this import. Put a `go.mod` next to the IaC file (or
in a parent directory):

```
module railway-iac

go 1.22

require github.com/railwayapp/railway-go-iac v0.2.0
```

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

The import path ends in `/iac` because the module path contains hyphens (not a
valid Go package name). Config as Code migration (`railway config migrate --lang go`)
emits the same import.

## Release (maintainers)

1. Merge to `main`.
2. Tag the release commit (version is the tag; `go.mod` has no version field):

   ```bash
   git tag v0.2.0
   git push origin v0.2.0
   ```

3. The tag workflow runs tests, creates a GitHub Release, and warms
   `proxy.golang.org` / `sum.golang.org`.
4. Docs: https://pkg.go.dev/github.com/railwayapp/railway-go-iac/iac

Do not tag off a commit that is not `main`. First public version should be
`v0.2.0` (authoring-parity). Stay on `v0.x` until the graph contract is stable;
a future `v2` would require a `/v2` module path.

Last resort only: declare `const Partial = "api"` in `railway.go` (same role
as `export const partial` in TypeScript). Do not rename a partial after apply.

See https://docs.railway.com/infrastructure-as-code
