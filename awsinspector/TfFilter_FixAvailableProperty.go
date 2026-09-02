package awsinspector


// Experimental.
type TfFilter_FixAvailableProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#comparison TfFilter#comparison}.
	// Experimental.
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#value TfFilter#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

