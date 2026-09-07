package billing


// Experimental.
type AwsView_DimensionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/billing_view#key AwsView#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/billing_view#values AwsView#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

