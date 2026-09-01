package awselb


// Experimental.
type AwsLbListenerRule_HostHeaderRewriteConfigProperty struct {
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#rewrite AwsLbListenerRule#rewrite}
	// Experimental.
	Rewrite *AwsLbListenerRule_TransformHostHeaderRewriteConfigRewriteProperty `field:"optional" json:"rewrite" yaml:"rewrite"`
}

