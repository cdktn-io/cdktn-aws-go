package awskendra


// Experimental.
type AwsKendraIndex_UserTokenConfigurationsProperty struct {
	// json_token_type_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#json_token_type_configuration AwsKendraIndex#json_token_type_configuration}
	// Experimental.
	JsonTokenTypeConfiguration *AwsKendraIndex_JsonTokenTypeConfigurationProperty `field:"optional" json:"jsonTokenTypeConfiguration" yaml:"jsonTokenTypeConfiguration"`
	// jwt_token_type_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#jwt_token_type_configuration AwsKendraIndex#jwt_token_type_configuration}
	// Experimental.
	JwtTokenTypeConfiguration *AwsKendraIndex_JwtTokenTypeConfigurationProperty `field:"optional" json:"jwtTokenTypeConfiguration" yaml:"jwtTokenTypeConfiguration"`
}

