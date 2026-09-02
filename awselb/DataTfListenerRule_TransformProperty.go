package awselb


// Experimental.
type DataTfListenerRule_TransformProperty struct {
	// host_header_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#host_header_rewrite_config DataTfListenerRule#host_header_rewrite_config}
	// Experimental.
	HostHeaderRewriteConfig interface{} `field:"optional" json:"hostHeaderRewriteConfig" yaml:"hostHeaderRewriteConfig"`
	// url_rewrite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#url_rewrite_config DataTfListenerRule#url_rewrite_config}
	// Experimental.
	UrlRewriteConfig interface{} `field:"optional" json:"urlRewriteConfig" yaml:"urlRewriteConfig"`
}

