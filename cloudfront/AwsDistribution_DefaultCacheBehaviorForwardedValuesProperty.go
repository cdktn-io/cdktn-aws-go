package cloudfront


// Experimental.
type AwsDistribution_DefaultCacheBehaviorForwardedValuesProperty struct {
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#cookies AwsDistribution#cookies}
	// Experimental.
	Cookies *AwsDistribution_DefaultCacheBehaviorForwardedValuesCookiesProperty `field:"required" json:"cookies" yaml:"cookies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#query_string AwsDistribution#query_string}.
	// Experimental.
	QueryString interface{} `field:"required" json:"queryString" yaml:"queryString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#headers AwsDistribution#headers}.
	// Experimental.
	Headers *[]*string `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#query_string_cache_keys AwsDistribution#query_string_cache_keys}.
	// Experimental.
	QueryStringCacheKeys *[]*string `field:"optional" json:"queryStringCacheKeys" yaml:"queryStringCacheKeys"`
}

