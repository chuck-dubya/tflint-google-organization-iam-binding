package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleOrganizationIamBindingRule checks for the use of google_organization_iam_binding in Terraform files
type GoogleOrganizationIamBindingRule struct {
	tflint.DefaultRule
}

// NewGoogleOrganizationIamBindingRule creates a new instance of the rule
func NewGoogleOrganizationIamBindingRule() *GoogleOrganizationIamBindingRule {
	return &GoogleOrganizationIamBindingRule{}
}

// Name returns the name of the rule
func (r *GoogleOrganizationIamBindingRule) Name() string {
	return "google_organization_iam_binding"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleOrganizationIamBindingRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleOrganizationIamBindingRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the reference link for the rule
func (r *GoogleOrganizationIamBindingRule) Link() string {
	return "https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_organization_iam_binding"
}

// Check checks for google_organization_iam_binding resource
func (r *GoogleOrganizationIamBindingRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("google_organization_iam_binding", nil, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		runner.EmitIssue(r, "google_organization_iam_binding usage detected", resource.DefRange)
	}

	return nil
}
