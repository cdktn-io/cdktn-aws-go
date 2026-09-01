package awselb


// Experimental.
type AwsLbListenerRule_UrlRewriteConfigProperty struct {
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#rewrite AwsLbListenerRule#rewrite}
	// Experimental.
	Rewrite *AwsLbListenerRule_TransformUrlRewriteConfigRewriteProperty `field:"optional" json:"rewrite" yaml:"rewrite"`
}

