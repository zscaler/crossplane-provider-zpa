/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"
	"github.com/zscaler/provider-zpa/config/zpaapplicationserver"
	"github.com/zscaler/provider-zpa/config/zpasegmentgroup"
)

const (
	resourcePrefix = "zpa"
	modulePath     = "github.com/zscaler/provider-zpa"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithShortName("zpa"),
		ujconfig.WithRootGroup("zpa.crossplane.io"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		zpaapplicationserver.Configure,
		zpasegmentgroup.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}