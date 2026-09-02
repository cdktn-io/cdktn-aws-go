package awscloudfront


// Experimental.
type TfCachePolicy_QueryStringsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#query_string_behavior TfCachePolicy#query_string_behavior}.
	// Experimental.
	QueryStringBehavior *string `field:"required" json:"queryStringBehavior" yaml:"queryStringBehavior"`
	// query_strings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#query_strings TfCachePolicy#query_strings}
	// Experimental.
	QueryStrings *TfCachePolicy_QueryStringsProperty `field:"optional" json:"queryStrings" yaml:"queryStrings"`
}

