package awscloudfront


// Experimental.
type AwsCloudfrontConnectionFunction_ConnectionFunctionConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_connection_function#comment AwsCloudfrontConnectionFunction#comment}.
	// Experimental.
	Comment *string `field:"required" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_connection_function#runtime AwsCloudfrontConnectionFunction#runtime}.
	// Experimental.
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// key_value_store_association block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_connection_function#key_value_store_association AwsCloudfrontConnectionFunction#key_value_store_association}
	// Experimental.
	KeyValueStoreAssociation interface{} `field:"optional" json:"keyValueStoreAssociation" yaml:"keyValueStoreAssociation"`
}

