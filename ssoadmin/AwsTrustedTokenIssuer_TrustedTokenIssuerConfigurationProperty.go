package ssoadmin


// Experimental.
type AwsTrustedTokenIssuer_TrustedTokenIssuerConfigurationProperty struct {
	// oidc_jwt_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_trusted_token_issuer#oidc_jwt_configuration AwsTrustedTokenIssuer#oidc_jwt_configuration}
	// Experimental.
	OidcJwtConfiguration interface{} `field:"optional" json:"oidcJwtConfiguration" yaml:"oidcJwtConfiguration"`
}

