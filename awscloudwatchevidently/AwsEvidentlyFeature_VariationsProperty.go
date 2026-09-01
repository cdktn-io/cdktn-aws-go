package awscloudwatchevidently


// Experimental.
type AwsEvidentlyFeature_VariationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_feature#name AwsEvidentlyFeature#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_feature#value AwsEvidentlyFeature#value}
	// Experimental.
	Value *AwsEvidentlyFeature_ValueProperty `field:"required" json:"value" yaml:"value"`
}

