package awscloudfront


// Experimental.
type AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty struct {
	// cookies_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#cookies_config AwsCloudfrontCachePolicy#cookies_config}
	// Experimental.
	CookiesConfig *AwsCloudfrontCachePolicy_CookiesConfigProperty `field:"required" json:"cookiesConfig" yaml:"cookiesConfig"`
	// headers_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#headers_config AwsCloudfrontCachePolicy#headers_config}
	// Experimental.
	HeadersConfig *AwsCloudfrontCachePolicy_HeadersConfigProperty `field:"required" json:"headersConfig" yaml:"headersConfig"`
	// query_strings_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#query_strings_config AwsCloudfrontCachePolicy#query_strings_config}
	// Experimental.
	QueryStringsConfig *AwsCloudfrontCachePolicy_QueryStringsConfigProperty `field:"required" json:"queryStringsConfig" yaml:"queryStringsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#enable_accept_encoding_brotli AwsCloudfrontCachePolicy#enable_accept_encoding_brotli}.
	// Experimental.
	EnableAcceptEncodingBrotli interface{} `field:"optional" json:"enableAcceptEncodingBrotli" yaml:"enableAcceptEncodingBrotli"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#enable_accept_encoding_gzip AwsCloudfrontCachePolicy#enable_accept_encoding_gzip}.
	// Experimental.
	EnableAcceptEncodingGzip interface{} `field:"optional" json:"enableAcceptEncodingGzip" yaml:"enableAcceptEncodingGzip"`
}

