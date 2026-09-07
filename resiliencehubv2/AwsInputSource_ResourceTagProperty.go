package resiliencehubv2


// Experimental.
type AwsInputSource_ResourceTagProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_input_source#key AwsInputSource#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_input_source#values AwsInputSource#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

