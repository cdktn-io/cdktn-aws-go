package inspector


// Experimental.
type AwsFilter_PortRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#begin_inclusive AwsFilter#begin_inclusive}.
	// Experimental.
	BeginInclusive *float64 `field:"required" json:"beginInclusive" yaml:"beginInclusive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#end_inclusive AwsFilter#end_inclusive}.
	// Experimental.
	EndInclusive *float64 `field:"required" json:"endInclusive" yaml:"endInclusive"`
}

