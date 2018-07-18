package main

import (
	"github.com/dmacvicar/terraform-provider-salt/salt"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: salt.Provider,
	})
}
