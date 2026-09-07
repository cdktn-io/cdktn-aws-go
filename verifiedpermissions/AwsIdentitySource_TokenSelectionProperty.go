package verifiedpermissions


// Experimental.
type AwsIdentitySource_TokenSelectionProperty struct {
	// access_token_only block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#access_token_only AwsIdentitySource#access_token_only}
	// Experimental.
	AccessTokenOnly interface{} `field:"optional" json:"accessTokenOnly" yaml:"accessTokenOnly"`
	// identity_token_only block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_identity_source#identity_token_only AwsIdentitySource#identity_token_only}
	// Experimental.
	IdentityTokenOnly interface{} `field:"optional" json:"identityTokenOnly" yaml:"identityTokenOnly"`
}

