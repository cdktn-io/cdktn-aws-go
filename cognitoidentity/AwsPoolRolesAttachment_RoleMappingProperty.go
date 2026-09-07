package cognitoidentity


// Experimental.
type AwsPoolRolesAttachment_RoleMappingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_identity_pool_roles_attachment#identity_provider AwsPoolRolesAttachment#identity_provider}.
	// Experimental.
	IdentityProvider *string `field:"required" json:"identityProvider" yaml:"identityProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_identity_pool_roles_attachment#type AwsPoolRolesAttachment#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_identity_pool_roles_attachment#ambiguous_role_resolution AwsPoolRolesAttachment#ambiguous_role_resolution}.
	// Experimental.
	AmbiguousRoleResolution *string `field:"optional" json:"ambiguousRoleResolution" yaml:"ambiguousRoleResolution"`
	// mapping_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_identity_pool_roles_attachment#mapping_rule AwsPoolRolesAttachment#mapping_rule}
	// Experimental.
	MappingRule interface{} `field:"optional" json:"mappingRule" yaml:"mappingRule"`
}

