package awseventbridge


// Experimental.
type TfConnection_ApiKeyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#key TfConnection#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#value TfConnection#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

