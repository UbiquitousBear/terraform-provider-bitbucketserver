package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/ubiquitousbear/terraform-provider-bitbucketserver/bitbucket"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: bitbucket.Provider})
}
