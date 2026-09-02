package awselb


// Experimental.
type TfAlbListenerRule_JwtValidationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#issuer TfAlbListenerRule#issuer}.
	// Experimental.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#jwks_endpoint TfAlbListenerRule#jwks_endpoint}.
	// Experimental.
	JwksEndpoint *string `field:"required" json:"jwksEndpoint" yaml:"jwksEndpoint"`
	// additional_claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#additional_claim TfAlbListenerRule#additional_claim}
	// Experimental.
	AdditionalClaim interface{} `field:"optional" json:"additionalClaim" yaml:"additionalClaim"`
}

