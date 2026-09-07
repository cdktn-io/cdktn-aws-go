package config


// Experimental.
type AwsOrganizationCustomRule_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_rule#create AwsOrganizationCustomRule#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_rule#delete AwsOrganizationCustomRule#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_custom_rule#update AwsOrganizationCustomRule#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

