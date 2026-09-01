package awselb


// Experimental.
type AwsAlbListenerRule_TransformHostHeaderRewriteConfigRewriteProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#regex AwsAlbListenerRule#regex}.
	// Experimental.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#replace AwsAlbListenerRule#replace}.
	// Experimental.
	Replace *string `field:"required" json:"replace" yaml:"replace"`
}

