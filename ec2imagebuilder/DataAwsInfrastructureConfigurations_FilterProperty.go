package ec2imagebuilder


// Experimental.
type DataAwsInfrastructureConfigurations_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_infrastructure_configurations#name DataAwsInfrastructureConfigurations#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_infrastructure_configurations#values DataAwsInfrastructureConfigurations#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

