package cloudscaleobjectsuser

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_objects_user", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		if s, ok := r.TerraformResource.Schema["keys"]; ok {
			s.Optional = false
			s.Computed = true
			s.Sensitive = true
		}
	})
}
