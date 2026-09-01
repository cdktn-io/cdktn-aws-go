package awsquicksight


// Experimental.
type AwsQuicksightAnalysis_IntegerParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_analysis#name AwsQuicksightAnalysis#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_analysis#values AwsQuicksightAnalysis#values}.
	// Experimental.
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

