package config


// Experimental.
type AwsOrganizationManagedRule_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#create AwsOrganizationManagedRule#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#delete AwsOrganizationManagedRule#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_organization_managed_rule#update AwsOrganizationManagedRule#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

