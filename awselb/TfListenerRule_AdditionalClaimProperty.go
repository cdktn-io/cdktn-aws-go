package awselb


// Experimental.
type TfListenerRule_AdditionalClaimProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#format TfListenerRule#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#name TfListenerRule#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#values TfListenerRule#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

