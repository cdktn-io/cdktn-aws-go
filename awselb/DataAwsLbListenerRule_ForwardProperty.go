package awselb


// Experimental.
type DataAwsLbListenerRule_ForwardProperty struct {
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#stickiness DataAwsLbListenerRule#stickiness}
	// Experimental.
	Stickiness interface{} `field:"optional" json:"stickiness" yaml:"stickiness"`
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#target_group DataAwsLbListenerRule#target_group}
	// Experimental.
	TargetGroup interface{} `field:"optional" json:"targetGroup" yaml:"targetGroup"`
}

