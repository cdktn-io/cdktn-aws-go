package awselb


// Experimental.
type DataTfListenerRule_ForwardProperty struct {
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#stickiness DataTfListenerRule#stickiness}
	// Experimental.
	Stickiness interface{} `field:"optional" json:"stickiness" yaml:"stickiness"`
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#target_group DataTfListenerRule#target_group}
	// Experimental.
	TargetGroup interface{} `field:"optional" json:"targetGroup" yaml:"targetGroup"`
}

