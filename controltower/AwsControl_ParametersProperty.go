package controltower


// Experimental.
type AwsControl_ParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/controltower_control#key AwsControl#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/controltower_control#value AwsControl#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

