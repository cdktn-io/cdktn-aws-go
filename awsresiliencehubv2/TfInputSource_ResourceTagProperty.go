package awsresiliencehubv2


// Experimental.
type TfInputSource_ResourceTagProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_input_source#key TfInputSource#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_input_source#values TfInputSource#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

