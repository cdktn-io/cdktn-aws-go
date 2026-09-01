package awselb


// Experimental.
type AwsAlbListenerRule_UrlRewriteConfigProperty struct {
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#rewrite AwsAlbListenerRule#rewrite}
	// Experimental.
	Rewrite *AwsAlbListenerRule_TransformUrlRewriteConfigRewriteProperty `field:"optional" json:"rewrite" yaml:"rewrite"`
}

