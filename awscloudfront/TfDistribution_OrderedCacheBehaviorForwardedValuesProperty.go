package awscloudfront


// Experimental.
type TfDistribution_OrderedCacheBehaviorForwardedValuesProperty struct {
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#cookies TfDistribution#cookies}
	// Experimental.
	Cookies *TfDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty `field:"required" json:"cookies" yaml:"cookies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#query_string TfDistribution#query_string}.
	// Experimental.
	QueryString interface{} `field:"required" json:"queryString" yaml:"queryString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#headers TfDistribution#headers}.
	// Experimental.
	Headers *[]*string `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#query_string_cache_keys TfDistribution#query_string_cache_keys}.
	// Experimental.
	QueryStringCacheKeys *[]*string `field:"optional" json:"queryStringCacheKeys" yaml:"queryStringCacheKeys"`
}

