package awselb


// Experimental.
type AwsLbListenerRule_ForwardProperty struct {
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#target_group AwsLbListenerRule#target_group}
	// Experimental.
	TargetGroup interface{} `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#stickiness AwsLbListenerRule#stickiness}
	// Experimental.
	Stickiness *AwsLbListenerRule_StickinessProperty `field:"optional" json:"stickiness" yaml:"stickiness"`
}

