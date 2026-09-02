package awsverifiedpermissions


// Experimental.
type TfIdentitySource_ConfigurationProperty struct {
	// cognito_user_pool_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#cognito_user_pool_configuration TfIdentitySource#cognito_user_pool_configuration}
	// Experimental.
	CognitoUserPoolConfiguration interface{} `field:"optional" json:"cognitoUserPoolConfiguration" yaml:"cognitoUserPoolConfiguration"`
	// open_id_connect_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#open_id_connect_configuration TfIdentitySource#open_id_connect_configuration}
	// Experimental.
	OpenIdConnectConfiguration interface{} `field:"optional" json:"openIdConnectConfiguration" yaml:"openIdConnectConfiguration"`
}

