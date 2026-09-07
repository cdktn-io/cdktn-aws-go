package elb


// Experimental.
type AwsListenerRule_UrlRewriteConfigProperty struct {
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#rewrite AwsListenerRule#rewrite}
	// Experimental.
	Rewrite *AwsListenerRule_TransformUrlRewriteConfigRewriteProperty `field:"optional" json:"rewrite" yaml:"rewrite"`
}

