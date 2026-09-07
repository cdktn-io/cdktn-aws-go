package appfabric


// Experimental.
type AwsAppAuthorization_CredentialProperty struct {
	// api_key_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_app_authorization#api_key_credential AwsAppAuthorization#api_key_credential}
	// Experimental.
	ApiKeyCredential interface{} `field:"optional" json:"apiKeyCredential" yaml:"apiKeyCredential"`
	// oauth2_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_app_authorization#oauth2_credential AwsAppAuthorization#oauth2_credential}
	// Experimental.
	Oauth2Credential interface{} `field:"optional" json:"oauth2Credential" yaml:"oauth2Credential"`
}

