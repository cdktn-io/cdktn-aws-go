package awselb


// Experimental.
type TfListenerRule_JwtValidationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#issuer TfListenerRule#issuer}.
	// Experimental.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#jwks_endpoint TfListenerRule#jwks_endpoint}.
	// Experimental.
	JwksEndpoint *string `field:"required" json:"jwksEndpoint" yaml:"jwksEndpoint"`
	// additional_claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener_rule#additional_claim TfListenerRule#additional_claim}
	// Experimental.
	AdditionalClaim interface{} `field:"optional" json:"additionalClaim" yaml:"additionalClaim"`
}

