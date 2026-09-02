package awssagemakerai


// Experimental.
type TfWorkteam_MemberDefinitionProperty struct {
	// cognito_member_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workteam#cognito_member_definition TfWorkteam#cognito_member_definition}
	// Experimental.
	CognitoMemberDefinition *TfWorkteam_CognitoMemberDefinitionProperty `field:"optional" json:"cognitoMemberDefinition" yaml:"cognitoMemberDefinition"`
	// oidc_member_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workteam#oidc_member_definition TfWorkteam#oidc_member_definition}
	// Experimental.
	OidcMemberDefinition *TfWorkteam_OidcMemberDefinitionProperty `field:"optional" json:"oidcMemberDefinition" yaml:"oidcMemberDefinition"`
}

