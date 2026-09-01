package awsglue


// Experimental.
type AwsGlueConnection_Oauth2PropertiesProperty struct {
	// authorization_code_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#authorization_code_properties AwsGlueConnection#authorization_code_properties}
	// Experimental.
	AuthorizationCodeProperties *AwsGlueConnection_AuthorizationCodePropertiesProperty `field:"optional" json:"authorizationCodeProperties" yaml:"authorizationCodeProperties"`
	// oauth2_client_application block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#oauth2_client_application AwsGlueConnection#oauth2_client_application}
	// Experimental.
	Oauth2ClientApplication *AwsGlueConnection_Oauth2ClientApplicationProperty `field:"optional" json:"oauth2ClientApplication" yaml:"oauth2ClientApplication"`
	// oauth2_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#oauth2_credentials AwsGlueConnection#oauth2_credentials}
	// Experimental.
	Oauth2Credentials *AwsGlueConnection_Oauth2CredentialsProperty `field:"optional" json:"oauth2Credentials" yaml:"oauth2Credentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#oauth2_grant_type AwsGlueConnection#oauth2_grant_type}.
	// Experimental.
	Oauth2GrantType *string `field:"optional" json:"oauth2GrantType" yaml:"oauth2GrantType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#token_url AwsGlueConnection#token_url}.
	// Experimental.
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#token_url_parameters_map AwsGlueConnection#token_url_parameters_map}.
	// Experimental.
	TokenUrlParametersMap *map[string]*string `field:"optional" json:"tokenUrlParametersMap" yaml:"tokenUrlParametersMap"`
}

