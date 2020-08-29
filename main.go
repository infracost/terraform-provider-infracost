package main

import (
	"terraform-provider-infracost/infracost"

	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: infracost.Provider})
}
