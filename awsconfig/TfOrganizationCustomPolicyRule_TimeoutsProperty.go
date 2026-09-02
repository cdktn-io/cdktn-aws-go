package awsconfig


// Experimental.
type TfOrganizationCustomPolicyRule_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#create TfOrganizationCustomPolicyRule#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#delete TfOrganizationCustomPolicyRule#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#update TfOrganizationCustomPolicyRule#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

