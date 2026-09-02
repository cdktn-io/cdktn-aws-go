package awsec2imagebuilder


// Experimental.
type DataTfDistributionConfigurations_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_distribution_configurations#name DataTfDistributionConfigurations#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_distribution_configurations#values DataTfDistributionConfigurations#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

