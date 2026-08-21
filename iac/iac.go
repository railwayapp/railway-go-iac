// Package iac is a thin Railway Infrastructure as Code authoring mirror.
// It builds the same conceptual RailwayGraph as railway/iac (TypeScript).
// Plan/apply stay in the CLI; this package has no Config as Code knowledge.
// Multi-repo: declare `const Partial = "api"` in railway.go (same role as
// `export const partial` in TypeScript).
package iac

type ServiceConfig map[string]any

type Service struct {
	Name   string
	Config ServiceConfig
}

func ServiceNamed(name string, config ServiceConfig) Service {
	if config == nil {
		config = ServiceConfig{}
	}
	return Service{Name: name, Config: config}
}

type Project struct {
	Name      string
	Resources []any
}

func ProjectNamed(name string, resources []any) Project {
	return Project{Name: name, Resources: resources}
}

// Graph is a JSON-serializable project definition for CLI evaluation.
func (p Project) Graph() map[string]any {
	resources := make([]any, 0, len(p.Resources))
	for _, r := range p.Resources {
		switch v := r.(type) {
		case Service:
			node := map[string]any{"type": "service", "name": v.Name}
			for k, val := range v.Config {
				node[k] = val
			}
			resources = append(resources, node)
		default:
			resources = append(resources, r)
		}
	}
	return map[string]any{
		"name":      p.Name,
		"resources": resources,
	}
}
