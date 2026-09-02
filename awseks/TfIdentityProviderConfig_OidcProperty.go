package awseks


// Experimental.
type TfIdentityProviderConfig_OidcProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#client_id TfIdentityProviderConfig#client_id}.
	// Experimental.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#identity_provider_config_name TfIdentityProviderConfig#identity_provider_config_name}.
	// Experimental.
	IdentityProviderConfigName *string `field:"required" json:"identityProviderConfigName" yaml:"identityProviderConfigName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#issuer_url TfIdentityProviderConfig#issuer_url}.
	// Experimental.
	IssuerUrl *string `field:"required" json:"issuerUrl" yaml:"issuerUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#groups_claim TfIdentityProviderConfig#groups_claim}.
	// Experimental.
	GroupsClaim *string `field:"optional" json:"groupsClaim" yaml:"groupsClaim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#groups_prefix TfIdentityProviderConfig#groups_prefix}.
	// Experimental.
	GroupsPrefix *string `field:"optional" json:"groupsPrefix" yaml:"groupsPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#required_claims TfIdentityProviderConfig#required_claims}.
	// Experimental.
	RequiredClaims *map[string]*string `field:"optional" json:"requiredClaims" yaml:"requiredClaims"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#username_claim TfIdentityProviderConfig#username_claim}.
	// Experimental.
	UsernameClaim *string `field:"optional" json:"usernameClaim" yaml:"usernameClaim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#username_prefix TfIdentityProviderConfig#username_prefix}.
	// Experimental.
	UsernamePrefix *string `field:"optional" json:"usernamePrefix" yaml:"usernamePrefix"`
}

