package elb


// Experimental.
type AwsListener_ForwardProperty struct {
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#target_group AwsListener#target_group}
	// Experimental.
	TargetGroup interface{} `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#stickiness AwsListener#stickiness}
	// Experimental.
	Stickiness *AwsListener_StickinessProperty `field:"optional" json:"stickiness" yaml:"stickiness"`
}

