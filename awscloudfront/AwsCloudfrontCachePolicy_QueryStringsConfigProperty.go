package awscloudfront


// Experimental.
type AwsCloudfrontCachePolicy_QueryStringsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#query_string_behavior AwsCloudfrontCachePolicy#query_string_behavior}.
	// Experimental.
	QueryStringBehavior *string `field:"required" json:"queryStringBehavior" yaml:"queryStringBehavior"`
	// query_strings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#query_strings AwsCloudfrontCachePolicy#query_strings}
	// Experimental.
	QueryStrings *AwsCloudfrontCachePolicy_QueryStringsProperty `field:"optional" json:"queryStrings" yaml:"queryStrings"`
}

