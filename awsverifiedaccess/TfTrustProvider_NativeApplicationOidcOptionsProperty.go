package awsverifiedaccess


// Experimental.
type TfTrustProvider_NativeApplicationOidcOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#client_secret TfTrustProvider#client_secret}.
	// Experimental.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#authorization_endpoint TfTrustProvider#authorization_endpoint}.
	// Experimental.
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#client_id TfTrustProvider#client_id}.
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#issuer TfTrustProvider#issuer}.
	// Experimental.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#public_signing_key_endpoint TfTrustProvider#public_signing_key_endpoint}.
	// Experimental.
	PublicSigningKeyEndpoint *string `field:"optional" json:"publicSigningKeyEndpoint" yaml:"publicSigningKeyEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#scope TfTrustProvider#scope}.
	// Experimental.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#token_endpoint TfTrustProvider#token_endpoint}.
	// Experimental.
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#user_info_endpoint TfTrustProvider#user_info_endpoint}.
	// Experimental.
	UserInfoEndpoint *string `field:"optional" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
}

