package awselb


// Experimental.
type TfListener_JwtValidationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#issuer TfListener#issuer}.
	// Experimental.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#jwks_endpoint TfListener#jwks_endpoint}.
	// Experimental.
	JwksEndpoint *string `field:"required" json:"jwksEndpoint" yaml:"jwksEndpoint"`
	// additional_claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener#additional_claim TfListener#additional_claim}
	// Experimental.
	AdditionalClaim interface{} `field:"optional" json:"additionalClaim" yaml:"additionalClaim"`
}

