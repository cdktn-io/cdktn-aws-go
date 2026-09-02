package awselb


// Experimental.
type TfAlbListener_JwtValidationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#issuer TfAlbListener#issuer}.
	// Experimental.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#jwks_endpoint TfAlbListener#jwks_endpoint}.
	// Experimental.
	JwksEndpoint *string `field:"required" json:"jwksEndpoint" yaml:"jwksEndpoint"`
	// additional_claim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#additional_claim TfAlbListener#additional_claim}
	// Experimental.
	AdditionalClaim interface{} `field:"optional" json:"additionalClaim" yaml:"additionalClaim"`
}

