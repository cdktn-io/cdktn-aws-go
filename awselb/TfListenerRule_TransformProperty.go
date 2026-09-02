package awselb


// Experimental.
type TfListenerRule_TransformProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#type TfListenerRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// host_header_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#host_header_rewrite_config TfListenerRule#host_header_rewrite_config}
	// Experimental.
	HostHeaderRewriteConfig *TfListenerRule_HostHeaderRewriteConfigProperty `field:"optional" json:"hostHeaderRewriteConfig" yaml:"hostHeaderRewriteConfig"`
	// url_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#url_rewrite_config TfListenerRule#url_rewrite_config}
	// Experimental.
	UrlRewriteConfig *TfListenerRule_UrlRewriteConfigProperty `field:"optional" json:"urlRewriteConfig" yaml:"urlRewriteConfig"`
}

