package ec2imagebuilder


// Experimental.
type AwsDistributionConfiguration_LaunchPermissionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#organizational_unit_arns AwsDistributionConfiguration#organizational_unit_arns}.
	// Experimental.
	OrganizationalUnitArns *[]*string `field:"optional" json:"organizationalUnitArns" yaml:"organizationalUnitArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#organization_arns AwsDistributionConfiguration#organization_arns}.
	// Experimental.
	OrganizationArns *[]*string `field:"optional" json:"organizationArns" yaml:"organizationArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#user_groups AwsDistributionConfiguration#user_groups}.
	// Experimental.
	UserGroups *[]*string `field:"optional" json:"userGroups" yaml:"userGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#user_ids AwsDistributionConfiguration#user_ids}.
	// Experimental.
	UserIds *[]*string `field:"optional" json:"userIds" yaml:"userIds"`
}

