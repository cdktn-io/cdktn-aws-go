package awssagemakerai


// Experimental.
type AwsSagemakerWorkteam_MemberDefinitionProperty struct {
	// cognito_member_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workteam#cognito_member_definition AwsSagemakerWorkteam#cognito_member_definition}
	// Experimental.
	CognitoMemberDefinition *AwsSagemakerWorkteam_CognitoMemberDefinitionProperty `field:"optional" json:"cognitoMemberDefinition" yaml:"cognitoMemberDefinition"`
	// oidc_member_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workteam#oidc_member_definition AwsSagemakerWorkteam#oidc_member_definition}
	// Experimental.
	OidcMemberDefinition *AwsSagemakerWorkteam_OidcMemberDefinitionProperty `field:"optional" json:"oidcMemberDefinition" yaml:"oidcMemberDefinition"`
}

