package cloudwatchevidently


// Experimental.
type AwsFeature_VariationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_feature#name AwsFeature#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_feature#value AwsFeature#value}
	// Experimental.
	Value *AwsFeature_ValueProperty `field:"required" json:"value" yaml:"value"`
}

