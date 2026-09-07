package elb


// Experimental.
type AwsAlbListenerRule_TransformProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#type AwsAlbListenerRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// host_header_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#host_header_rewrite_config AwsAlbListenerRule#host_header_rewrite_config}
	// Experimental.
	HostHeaderRewriteConfig *AwsAlbListenerRule_HostHeaderRewriteConfigProperty `field:"optional" json:"hostHeaderRewriteConfig" yaml:"hostHeaderRewriteConfig"`
	// url_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#url_rewrite_config AwsAlbListenerRule#url_rewrite_config}
	// Experimental.
	UrlRewriteConfig *AwsAlbListenerRule_UrlRewriteConfigProperty `field:"optional" json:"urlRewriteConfig" yaml:"urlRewriteConfig"`
}

