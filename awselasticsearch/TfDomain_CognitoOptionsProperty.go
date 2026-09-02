package awselasticsearch


// Experimental.
type TfDomain_CognitoOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#identity_pool_id TfDomain#identity_pool_id}.
	// Experimental.
	IdentityPoolId *string `field:"required" json:"identityPoolId" yaml:"identityPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#role_arn TfDomain#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#user_pool_id TfDomain#user_pool_id}.
	// Experimental.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticsearch_domain#enabled TfDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

