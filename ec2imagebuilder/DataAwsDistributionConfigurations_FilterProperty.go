package ec2imagebuilder


// Experimental.
type DataAwsDistributionConfigurations_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_distribution_configurations#name DataAwsDistributionConfigurations#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_distribution_configurations#values DataAwsDistributionConfigurations#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

