package elb


// Experimental.
type AwsListenerRule_TransformHostHeaderRewriteConfigRewriteProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#regex AwsListenerRule#regex}.
	// Experimental.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#replace AwsListenerRule#replace}.
	// Experimental.
	Replace *string `field:"required" json:"replace" yaml:"replace"`
}

