package awsinspector


// Experimental.
type TfFilter_InspectorScoreProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#lower_inclusive TfFilter#lower_inclusive}.
	// Experimental.
	LowerInclusive *float64 `field:"required" json:"lowerInclusive" yaml:"lowerInclusive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#upper_inclusive TfFilter#upper_inclusive}.
	// Experimental.
	UpperInclusive *float64 `field:"required" json:"upperInclusive" yaml:"upperInclusive"`
}

