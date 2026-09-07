package elb


// Experimental.
type AwsAlbListener_ForwardProperty struct {
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#target_group AwsAlbListener#target_group}
	// Experimental.
	TargetGroup interface{} `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#stickiness AwsAlbListener#stickiness}
	// Experimental.
	Stickiness *AwsAlbListener_StickinessProperty `field:"optional" json:"stickiness" yaml:"stickiness"`
}

