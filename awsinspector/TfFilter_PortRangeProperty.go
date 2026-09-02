package awsinspector


// Experimental.
type TfFilter_PortRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#begin_inclusive TfFilter#begin_inclusive}.
	// Experimental.
	BeginInclusive *float64 `field:"required" json:"beginInclusive" yaml:"beginInclusive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#end_inclusive TfFilter#end_inclusive}.
	// Experimental.
	EndInclusive *float64 `field:"required" json:"endInclusive" yaml:"endInclusive"`
}

