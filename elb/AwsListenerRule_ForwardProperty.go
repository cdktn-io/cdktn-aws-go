package elb


// Experimental.
type AwsListenerRule_ForwardProperty struct {
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#target_group AwsListenerRule#target_group}
	// Experimental.
	TargetGroup interface{} `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#stickiness AwsListenerRule#stickiness}
	// Experimental.
	Stickiness *AwsListenerRule_StickinessProperty `field:"optional" json:"stickiness" yaml:"stickiness"`
}

