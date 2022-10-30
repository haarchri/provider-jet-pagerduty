package maintenance

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_maintenance_window", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "maintenance"
		r.References = config.References{
			"services": {
				Type:              "github.com/crossplane-contrib/provider-pagerduty/apis/service/v1alpha1.Service",
				RefFieldName:      "ServiceRefs",
				SelectorFieldName: "ServiceSelector",
			},
		}
	})
}
