package awselb


// Experimental.
type TfListenerRule_ForwardProperty struct {
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#target_group TfListenerRule#target_group}
	// Experimental.
	TargetGroup interface{} `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#stickiness TfListenerRule#stickiness}
	// Experimental.
	Stickiness *TfListenerRule_StickinessProperty `field:"optional" json:"stickiness" yaml:"stickiness"`
}

