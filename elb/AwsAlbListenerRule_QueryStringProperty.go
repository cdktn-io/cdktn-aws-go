package elb


// Experimental.
type AwsAlbListenerRule_QueryStringProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#value AwsAlbListenerRule#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#key AwsAlbListenerRule#key}.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

