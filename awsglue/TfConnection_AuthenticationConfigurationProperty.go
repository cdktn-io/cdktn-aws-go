package awsglue


// Experimental.
type TfConnection_AuthenticationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#authentication_type TfConnection#authentication_type}.
	// Experimental.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// basic_authentication_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#basic_authentication_credentials TfConnection#basic_authentication_credentials}
	// Experimental.
	BasicAuthenticationCredentials *TfConnection_BasicAuthenticationCredentialsProperty `field:"optional" json:"basicAuthenticationCredentials" yaml:"basicAuthenticationCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#custom_authentication_credentials TfConnection#custom_authentication_credentials}.
	// Experimental.
	CustomAuthenticationCredentials *map[string]*string `field:"optional" json:"customAuthenticationCredentials" yaml:"customAuthenticationCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#kms_key_arn TfConnection#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// oauth2_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#oauth2_properties TfConnection#oauth2_properties}
	// Experimental.
	Oauth2Properties *TfConnection_Oauth2PropertiesProperty `field:"optional" json:"oauth2Properties" yaml:"oauth2Properties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#secret_arn TfConnection#secret_arn}.
	// Experimental.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

