package awselb


// Experimental.
type TfAlbListenerRule_TransformProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#type TfAlbListenerRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// host_header_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#host_header_rewrite_config TfAlbListenerRule#host_header_rewrite_config}
	// Experimental.
	HostHeaderRewriteConfig *TfAlbListenerRule_HostHeaderRewriteConfigProperty `field:"optional" json:"hostHeaderRewriteConfig" yaml:"hostHeaderRewriteConfig"`
	// url_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#url_rewrite_config TfAlbListenerRule#url_rewrite_config}
	// Experimental.
	UrlRewriteConfig *TfAlbListenerRule_UrlRewriteConfigProperty `field:"optional" json:"urlRewriteConfig" yaml:"urlRewriteConfig"`
}

