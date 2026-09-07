package cloudfront


// Experimental.
type AwsMultitenantDistribution_ParameterDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#name AwsMultitenantDistribution#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#definition AwsMultitenantDistribution#definition}
	// Experimental.
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
}

