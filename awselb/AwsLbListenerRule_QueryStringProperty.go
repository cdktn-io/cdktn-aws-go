package awselb


// Experimental.
type AwsLbListenerRule_QueryStringProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#value AwsLbListenerRule#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#key AwsLbListenerRule#key}.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

