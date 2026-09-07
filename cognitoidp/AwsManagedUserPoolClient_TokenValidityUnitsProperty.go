package cognitoidp


// Experimental.
type AwsManagedUserPoolClient_TokenValidityUnitsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_user_pool_client#access_token AwsManagedUserPoolClient#access_token}.
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_user_pool_client#id_token AwsManagedUserPoolClient#id_token}.
	// Experimental.
	IdToken *string `field:"optional" json:"idToken" yaml:"idToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_user_pool_client#refresh_token AwsManagedUserPoolClient#refresh_token}.
	// Experimental.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

