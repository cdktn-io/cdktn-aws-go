package awselb


// Experimental.
type TfAlbListenerRule_ForwardProperty struct {
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#target_group TfAlbListenerRule#target_group}
	// Experimental.
	TargetGroup interface{} `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#stickiness TfAlbListenerRule#stickiness}
	// Experimental.
	Stickiness *TfAlbListenerRule_StickinessProperty `field:"optional" json:"stickiness" yaml:"stickiness"`
}

