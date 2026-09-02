package awskendra


// Experimental.
type TfIndex_UserTokenConfigurationsProperty struct {
	// json_token_type_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#json_token_type_configuration TfIndex#json_token_type_configuration}
	// Experimental.
	JsonTokenTypeConfiguration *TfIndex_JsonTokenTypeConfigurationProperty `field:"optional" json:"jsonTokenTypeConfiguration" yaml:"jsonTokenTypeConfiguration"`
	// jwt_token_type_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#jwt_token_type_configuration TfIndex#jwt_token_type_configuration}
	// Experimental.
	JwtTokenTypeConfiguration *TfIndex_JwtTokenTypeConfigurationProperty `field:"optional" json:"jwtTokenTypeConfiguration" yaml:"jwtTokenTypeConfiguration"`
}

