package awseventbridge


// Experimental.
type TfConnection_AuthParametersInvocationHttpParametersBodyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#is_value_secret TfConnection#is_value_secret}.
	// Experimental.
	IsValueSecret interface{} `field:"optional" json:"isValueSecret" yaml:"isValueSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#key TfConnection#key}.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#value TfConnection#value}.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

