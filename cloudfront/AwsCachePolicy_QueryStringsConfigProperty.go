package cloudfront


// Experimental.
type AwsCachePolicy_QueryStringsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#query_string_behavior AwsCachePolicy#query_string_behavior}.
	// Experimental.
	QueryStringBehavior *string `field:"required" json:"queryStringBehavior" yaml:"queryStringBehavior"`
	// query_strings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#query_strings AwsCachePolicy#query_strings}
	// Experimental.
	QueryStrings *AwsCachePolicy_QueryStringsProperty `field:"optional" json:"queryStrings" yaml:"queryStrings"`
}

