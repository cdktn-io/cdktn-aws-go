package awsinspector


// Experimental.
type AwsInspector2Filter_EcrImageRegistryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#comparison AwsInspector2Filter#comparison}.
	// Experimental.
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#value AwsInspector2Filter#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

