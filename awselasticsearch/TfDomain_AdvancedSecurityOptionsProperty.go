package awselasticsearch


// Experimental.
type TfDomain_AdvancedSecurityOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#enabled TfDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#internal_user_database_enabled TfDomain#internal_user_database_enabled}.
	// Experimental.
	InternalUserDatabaseEnabled interface{} `field:"optional" json:"internalUserDatabaseEnabled" yaml:"internalUserDatabaseEnabled"`
	// master_user_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#master_user_options TfDomain#master_user_options}
	// Experimental.
	MasterUserOptions *TfDomain_MasterUserOptionsProperty `field:"optional" json:"masterUserOptions" yaml:"masterUserOptions"`
}

