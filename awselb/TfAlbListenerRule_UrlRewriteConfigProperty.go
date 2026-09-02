package awselb


// Experimental.
type TfAlbListenerRule_UrlRewriteConfigProperty struct {
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#rewrite TfAlbListenerRule#rewrite}
	// Experimental.
	Rewrite *TfAlbListenerRule_TransformUrlRewriteConfigRewriteProperty `field:"optional" json:"rewrite" yaml:"rewrite"`
}

