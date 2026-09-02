package awselb


// Experimental.
type TfListenerRule_TransformHostHeaderRewriteConfigRewriteProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#regex TfListenerRule#regex}.
	// Experimental.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#replace TfListenerRule#replace}.
	// Experimental.
	Replace *string `field:"required" json:"replace" yaml:"replace"`
}

