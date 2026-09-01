package awsssoadmin


// Experimental.
type AwsSsoadminTrustedTokenIssuer_OidcJwtConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_trusted_token_issuer#claim_attribute_path AwsSsoadminTrustedTokenIssuer#claim_attribute_path}.
	// Experimental.
	ClaimAttributePath *string `field:"required" json:"claimAttributePath" yaml:"claimAttributePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_trusted_token_issuer#identity_store_attribute_path AwsSsoadminTrustedTokenIssuer#identity_store_attribute_path}.
	// Experimental.
	IdentityStoreAttributePath *string `field:"required" json:"identityStoreAttributePath" yaml:"identityStoreAttributePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_trusted_token_issuer#issuer_url AwsSsoadminTrustedTokenIssuer#issuer_url}.
	// Experimental.
	IssuerUrl *string `field:"required" json:"issuerUrl" yaml:"issuerUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_trusted_token_issuer#jwks_retrieval_option AwsSsoadminTrustedTokenIssuer#jwks_retrieval_option}.
	// Experimental.
	JwksRetrievalOption *string `field:"required" json:"jwksRetrievalOption" yaml:"jwksRetrievalOption"`
}

