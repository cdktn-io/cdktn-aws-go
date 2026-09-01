package awsresiliencehubv2


// Experimental.
type AwsResiliencehubv2InputSource_ResourceTagProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_input_source#key AwsResiliencehubv2InputSource#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_input_source#values AwsResiliencehubv2InputSource#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

