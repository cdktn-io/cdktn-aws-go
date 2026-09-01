package awselb


// Experimental.
type AwsAlbListenerRule_ForwardProperty struct {
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#target_group AwsAlbListenerRule#target_group}
	// Experimental.
	TargetGroup interface{} `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#stickiness AwsAlbListenerRule#stickiness}
	// Experimental.
	Stickiness *AwsAlbListenerRule_StickinessProperty `field:"optional" json:"stickiness" yaml:"stickiness"`
}

