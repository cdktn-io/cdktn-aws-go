package awsbilling


// Experimental.
type AwsBillingView_TagsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/billing_view#key AwsBillingView#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/billing_view#values AwsBillingView#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

