package awselb


// Experimental.
type TfAlbListener_ForwardProperty struct {
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#target_group TfAlbListener#target_group}
	// Experimental.
	TargetGroup interface{} `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#stickiness TfAlbListener#stickiness}
	// Experimental.
	Stickiness *TfAlbListener_StickinessProperty `field:"optional" json:"stickiness" yaml:"stickiness"`
}

