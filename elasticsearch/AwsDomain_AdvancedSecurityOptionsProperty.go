package elasticsearch


// Experimental.
type AwsDomain_AdvancedSecurityOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#enabled AwsDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#internal_user_database_enabled AwsDomain#internal_user_database_enabled}.
	// Experimental.
	InternalUserDatabaseEnabled interface{} `field:"optional" json:"internalUserDatabaseEnabled" yaml:"internalUserDatabaseEnabled"`
	// master_user_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#master_user_options AwsDomain#master_user_options}
	// Experimental.
	MasterUserOptions *AwsDomain_MasterUserOptionsProperty `field:"optional" json:"masterUserOptions" yaml:"masterUserOptions"`
}

