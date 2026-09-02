package awsglue


// Experimental.
type TfConnection_Oauth2PropertiesProperty struct {
	// authorization_code_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#authorization_code_properties TfConnection#authorization_code_properties}
	// Experimental.
	AuthorizationCodeProperties *TfConnection_AuthorizationCodePropertiesProperty `field:"optional" json:"authorizationCodeProperties" yaml:"authorizationCodeProperties"`
	// oauth2_client_application block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#oauth2_client_application TfConnection#oauth2_client_application}
	// Experimental.
	Oauth2ClientApplication *TfConnection_Oauth2ClientApplicationProperty `field:"optional" json:"oauth2ClientApplication" yaml:"oauth2ClientApplication"`
	// oauth2_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#oauth2_credentials TfConnection#oauth2_credentials}
	// Experimental.
	Oauth2Credentials *TfConnection_Oauth2CredentialsProperty `field:"optional" json:"oauth2Credentials" yaml:"oauth2Credentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#oauth2_grant_type TfConnection#oauth2_grant_type}.
	// Experimental.
	Oauth2GrantType *string `field:"optional" json:"oauth2GrantType" yaml:"oauth2GrantType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#token_url TfConnection#token_url}.
	// Experimental.
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#token_url_parameters_map TfConnection#token_url_parameters_map}.
	// Experimental.
	TokenUrlParametersMap *map[string]*string `field:"optional" json:"tokenUrlParametersMap" yaml:"tokenUrlParametersMap"`
}

