package awscloudfront


// Experimental.
type AwsCloudfrontDistribution_OrderedCacheBehaviorForwardedValuesProperty struct {
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#cookies AwsCloudfrontDistribution#cookies}
	// Experimental.
	Cookies *AwsCloudfrontDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty `field:"required" json:"cookies" yaml:"cookies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#query_string AwsCloudfrontDistribution#query_string}.
	// Experimental.
	QueryString interface{} `field:"required" json:"queryString" yaml:"queryString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#headers AwsCloudfrontDistribution#headers}.
	// Experimental.
	Headers *[]*string `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#query_string_cache_keys AwsCloudfrontDistribution#query_string_cache_keys}.
	// Experimental.
	QueryStringCacheKeys *[]*string `field:"optional" json:"queryStringCacheKeys" yaml:"queryStringCacheKeys"`
}

