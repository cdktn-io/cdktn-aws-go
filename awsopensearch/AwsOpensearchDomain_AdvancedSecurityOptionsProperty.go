package awsopensearch


// Experimental.
type AwsOpensearchDomain_AdvancedSecurityOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#enabled AwsOpensearchDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#anonymous_auth_enabled AwsOpensearchDomain#anonymous_auth_enabled}.
	// Experimental.
	AnonymousAuthEnabled interface{} `field:"optional" json:"anonymousAuthEnabled" yaml:"anonymousAuthEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#internal_user_database_enabled AwsOpensearchDomain#internal_user_database_enabled}.
	// Experimental.
	InternalUserDatabaseEnabled interface{} `field:"optional" json:"internalUserDatabaseEnabled" yaml:"internalUserDatabaseEnabled"`
	// jwt_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#jwt_options AwsOpensearchDomain#jwt_options}
	// Experimental.
	JwtOptions *AwsOpensearchDomain_JwtOptionsProperty `field:"optional" json:"jwtOptions" yaml:"jwtOptions"`
	// master_user_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#master_user_options AwsOpensearchDomain#master_user_options}
	// Experimental.
	MasterUserOptions *AwsOpensearchDomain_MasterUserOptionsProperty `field:"optional" json:"masterUserOptions" yaml:"masterUserOptions"`
}

