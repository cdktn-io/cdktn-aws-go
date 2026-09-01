package awscognitoidentity


// Experimental.
type AwsCognitoIdentityPoolRolesAttachment_MappingRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_identity_pool_roles_attachment#claim AwsCognitoIdentityPoolRolesAttachment#claim}.
	// Experimental.
	Claim *string `field:"required" json:"claim" yaml:"claim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_identity_pool_roles_attachment#match_type AwsCognitoIdentityPoolRolesAttachment#match_type}.
	// Experimental.
	MatchType *string `field:"required" json:"matchType" yaml:"matchType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_identity_pool_roles_attachment#role_arn AwsCognitoIdentityPoolRolesAttachment#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_identity_pool_roles_attachment#value AwsCognitoIdentityPoolRolesAttachment#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

