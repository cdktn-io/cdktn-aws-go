package apigateway


// Experimental.
type AwsMethodSettings_SettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method_settings#cache_data_encrypted AwsMethodSettings#cache_data_encrypted}.
	// Experimental.
	CacheDataEncrypted interface{} `field:"optional" json:"cacheDataEncrypted" yaml:"cacheDataEncrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method_settings#cache_ttl_in_seconds AwsMethodSettings#cache_ttl_in_seconds}.
	// Experimental.
	CacheTtlInSeconds *float64 `field:"optional" json:"cacheTtlInSeconds" yaml:"cacheTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method_settings#caching_enabled AwsMethodSettings#caching_enabled}.
	// Experimental.
	CachingEnabled interface{} `field:"optional" json:"cachingEnabled" yaml:"cachingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method_settings#data_trace_enabled AwsMethodSettings#data_trace_enabled}.
	// Experimental.
	DataTraceEnabled interface{} `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method_settings#logging_level AwsMethodSettings#logging_level}.
	// Experimental.
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method_settings#metrics_enabled AwsMethodSettings#metrics_enabled}.
	// Experimental.
	MetricsEnabled interface{} `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method_settings#require_authorization_for_cache_control AwsMethodSettings#require_authorization_for_cache_control}.
	// Experimental.
	RequireAuthorizationForCacheControl interface{} `field:"optional" json:"requireAuthorizationForCacheControl" yaml:"requireAuthorizationForCacheControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method_settings#throttling_burst_limit AwsMethodSettings#throttling_burst_limit}.
	// Experimental.
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method_settings#throttling_rate_limit AwsMethodSettings#throttling_rate_limit}.
	// Experimental.
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_method_settings#unauthorized_cache_control_header_strategy AwsMethodSettings#unauthorized_cache_control_header_strategy}.
	// Experimental.
	UnauthorizedCacheControlHeaderStrategy *string `field:"optional" json:"unauthorizedCacheControlHeaderStrategy" yaml:"unauthorizedCacheControlHeaderStrategy"`
}

