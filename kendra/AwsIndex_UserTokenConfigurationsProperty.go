package kendra


// Experimental.
type AwsIndex_UserTokenConfigurationsProperty struct {
	// json_token_type_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#json_token_type_configuration AwsIndex#json_token_type_configuration}
	// Experimental.
	JsonTokenTypeConfiguration *AwsIndex_JsonTokenTypeConfigurationProperty `field:"optional" json:"jsonTokenTypeConfiguration" yaml:"jsonTokenTypeConfiguration"`
	// jwt_token_type_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#jwt_token_type_configuration AwsIndex#jwt_token_type_configuration}
	// Experimental.
	JwtTokenTypeConfiguration *AwsIndex_JwtTokenTypeConfigurationProperty `field:"optional" json:"jwtTokenTypeConfiguration" yaml:"jwtTokenTypeConfiguration"`
}

