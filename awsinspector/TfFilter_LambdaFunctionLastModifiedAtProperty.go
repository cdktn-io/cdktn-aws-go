package awsinspector


// Experimental.
type TfFilter_LambdaFunctionLastModifiedAtProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#end_inclusive TfFilter#end_inclusive}.
	// Experimental.
	EndInclusive *string `field:"optional" json:"endInclusive" yaml:"endInclusive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#start_inclusive TfFilter#start_inclusive}.
	// Experimental.
	StartInclusive *string `field:"optional" json:"startInclusive" yaml:"startInclusive"`
}

