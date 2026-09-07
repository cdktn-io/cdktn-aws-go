package elb


// Experimental.
type AwsAlbListenerRule_HostHeaderRewriteConfigProperty struct {
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#rewrite AwsAlbListenerRule#rewrite}
	// Experimental.
	Rewrite *AwsAlbListenerRule_TransformHostHeaderRewriteConfigRewriteProperty `field:"optional" json:"rewrite" yaml:"rewrite"`
}

