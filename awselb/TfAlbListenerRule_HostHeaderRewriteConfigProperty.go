package awselb


// Experimental.
type TfAlbListenerRule_HostHeaderRewriteConfigProperty struct {
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#rewrite TfAlbListenerRule#rewrite}
	// Experimental.
	Rewrite *TfAlbListenerRule_TransformHostHeaderRewriteConfigRewriteProperty `field:"optional" json:"rewrite" yaml:"rewrite"`
}

