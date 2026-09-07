package config


// Experimental.
type AwsOrganizationCustomPolicyRule_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#create AwsOrganizationCustomPolicyRule#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#delete AwsOrganizationCustomPolicyRule#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_policy_rule#update AwsOrganizationCustomPolicyRule#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

