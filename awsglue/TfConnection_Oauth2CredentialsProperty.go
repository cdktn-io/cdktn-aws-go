package awsglue


// Experimental.
type TfConnection_Oauth2CredentialsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#access_token TfConnection#access_token}.
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#jwt_token TfConnection#jwt_token}.
	// Experimental.
	JwtToken *string `field:"optional" json:"jwtToken" yaml:"jwtToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#refresh_token TfConnection#refresh_token}.
	// Experimental.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#user_managed_client_application_client_secret TfConnection#user_managed_client_application_client_secret}.
	// Experimental.
	UserManagedClientApplicationClientSecret *string `field:"optional" json:"userManagedClientApplicationClientSecret" yaml:"userManagedClientApplicationClientSecret"`
}

