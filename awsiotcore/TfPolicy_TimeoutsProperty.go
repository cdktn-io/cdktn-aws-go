package awsiotcore


// Experimental.
type TfPolicy_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_policy#delete TfPolicy#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_policy#update TfPolicy#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

