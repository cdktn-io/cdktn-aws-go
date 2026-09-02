package awsquicksight


// Experimental.
type TfAnalysis_StringParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_analysis#name TfAnalysis#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_analysis#values TfAnalysis#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

