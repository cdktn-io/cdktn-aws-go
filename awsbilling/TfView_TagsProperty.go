package awsbilling


// Experimental.
type TfView_TagsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/billing_view#key TfView#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/billing_view#values TfView#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

