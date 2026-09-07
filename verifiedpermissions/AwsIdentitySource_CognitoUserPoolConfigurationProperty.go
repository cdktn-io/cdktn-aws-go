package verifiedpermissions


// Experimental.
type AwsIdentitySource_CognitoUserPoolConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#user_pool_arn AwsIdentitySource#user_pool_arn}.
	// Experimental.
	UserPoolArn *string `field:"required" json:"userPoolArn" yaml:"userPoolArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#client_ids AwsIdentitySource#client_ids}.
	// Experimental.
	ClientIds *[]*string `field:"optional" json:"clientIds" yaml:"clientIds"`
	// group_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#group_configuration AwsIdentitySource#group_configuration}
	// Experimental.
	GroupConfiguration interface{} `field:"optional" json:"groupConfiguration" yaml:"groupConfiguration"`
}

