package zpaappconnectorgroup

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zpa_app_connector_group", func(r *config.Resource) {
		r.ShortGroup = "zpaappconnectorgroup"
	})
}
