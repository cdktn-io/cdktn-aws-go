package elb


// Experimental.
type AwsListenerRule_TransformProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#type AwsListenerRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// host_header_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#host_header_rewrite_config AwsListenerRule#host_header_rewrite_config}
	// Experimental.
	HostHeaderRewriteConfig *AwsListenerRule_HostHeaderRewriteConfigProperty `field:"optional" json:"hostHeaderRewriteConfig" yaml:"hostHeaderRewriteConfig"`
	// url_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#url_rewrite_config AwsListenerRule#url_rewrite_config}
	// Experimental.
	UrlRewriteConfig *AwsListenerRule_UrlRewriteConfigProperty `field:"optional" json:"urlRewriteConfig" yaml:"urlRewriteConfig"`
}

