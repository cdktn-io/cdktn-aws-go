package awseks


// Experimental.
type AwsEksIdentityProviderConfig_OidcProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#client_id AwsEksIdentityProviderConfig#client_id}.
	// Experimental.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#identity_provider_config_name AwsEksIdentityProviderConfig#identity_provider_config_name}.
	// Experimental.
	IdentityProviderConfigName *string `field:"required" json:"identityProviderConfigName" yaml:"identityProviderConfigName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#issuer_url AwsEksIdentityProviderConfig#issuer_url}.
	// Experimental.
	IssuerUrl *string `field:"required" json:"issuerUrl" yaml:"issuerUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#groups_claim AwsEksIdentityProviderConfig#groups_claim}.
	// Experimental.
	GroupsClaim *string `field:"optional" json:"groupsClaim" yaml:"groupsClaim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#groups_prefix AwsEksIdentityProviderConfig#groups_prefix}.
	// Experimental.
	GroupsPrefix *string `field:"optional" json:"groupsPrefix" yaml:"groupsPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#required_claims AwsEksIdentityProviderConfig#required_claims}.
	// Experimental.
	RequiredClaims *map[string]*string `field:"optional" json:"requiredClaims" yaml:"requiredClaims"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#username_claim AwsEksIdentityProviderConfig#username_claim}.
	// Experimental.
	UsernameClaim *string `field:"optional" json:"usernameClaim" yaml:"usernameClaim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_identity_provider_config#username_prefix AwsEksIdentityProviderConfig#username_prefix}.
	// Experimental.
	UsernamePrefix *string `field:"optional" json:"usernamePrefix" yaml:"usernamePrefix"`
}

