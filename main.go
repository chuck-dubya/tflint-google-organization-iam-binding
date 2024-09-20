package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/tflint-google-organization-iam-binding/rules" // Adjust this to match your local directory structure
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleFactories: []func() plugin.Rule{
			rules.NewGoogleOrganizationIamBindingRule,
		},
	})
}
