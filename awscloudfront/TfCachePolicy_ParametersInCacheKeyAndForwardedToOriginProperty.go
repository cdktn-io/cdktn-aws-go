package awscloudfront


// Experimental.
type TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty struct {
	// cookies_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#cookies_config TfCachePolicy#cookies_config}
	// Experimental.
	CookiesConfig *TfCachePolicy_CookiesConfigProperty `field:"required" json:"cookiesConfig" yaml:"cookiesConfig"`
	// headers_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#headers_config TfCachePolicy#headers_config}
	// Experimental.
	HeadersConfig *TfCachePolicy_HeadersConfigProperty `field:"required" json:"headersConfig" yaml:"headersConfig"`
	// query_strings_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#query_strings_config TfCachePolicy#query_strings_config}
	// Experimental.
	QueryStringsConfig *TfCachePolicy_QueryStringsConfigProperty `field:"required" json:"queryStringsConfig" yaml:"queryStringsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#enable_accept_encoding_brotli TfCachePolicy#enable_accept_encoding_brotli}.
	// Experimental.
	EnableAcceptEncodingBrotli interface{} `field:"optional" json:"enableAcceptEncodingBrotli" yaml:"enableAcceptEncodingBrotli"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#enable_accept_encoding_gzip TfCachePolicy#enable_accept_encoding_gzip}.
	// Experimental.
	EnableAcceptEncodingGzip interface{} `field:"optional" json:"enableAcceptEncodingGzip" yaml:"enableAcceptEncodingGzip"`
}

