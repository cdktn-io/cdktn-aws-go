package awslightsail


// Experimental.
type AwsLightsailDistribution_CacheBehaviorSettingsProperty struct {
	// The HTTP methods that are processed and forwarded to the distribution's origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#allowed_http_methods AwsLightsailDistribution#allowed_http_methods}
	// Experimental.
	AllowedHttpMethods *string `field:"optional" json:"allowedHttpMethods" yaml:"allowedHttpMethods"`
	// The HTTP method responses that are cached by your distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#cached_http_methods AwsLightsailDistribution#cached_http_methods}
	// Experimental.
	CachedHttpMethods *string `field:"optional" json:"cachedHttpMethods" yaml:"cachedHttpMethods"`
	// The default amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the content has been updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#default_ttl AwsLightsailDistribution#default_ttl}
	// Experimental.
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// forwarded_cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#forwarded_cookies AwsLightsailDistribution#forwarded_cookies}
	// Experimental.
	ForwardedCookies *AwsLightsailDistribution_ForwardedCookiesProperty `field:"optional" json:"forwardedCookies" yaml:"forwardedCookies"`
	// forwarded_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#forwarded_headers AwsLightsailDistribution#forwarded_headers}
	// Experimental.
	ForwardedHeaders *AwsLightsailDistribution_ForwardedHeadersProperty `field:"optional" json:"forwardedHeaders" yaml:"forwardedHeaders"`
	// forwarded_query_strings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#forwarded_query_strings AwsLightsailDistribution#forwarded_query_strings}
	// Experimental.
	ForwardedQueryStrings *AwsLightsailDistribution_ForwardedQueryStringsProperty `field:"optional" json:"forwardedQueryStrings" yaml:"forwardedQueryStrings"`
	// The maximum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#maximum_ttl AwsLightsailDistribution#maximum_ttl}
	// Experimental.
	MaximumTtl *float64 `field:"optional" json:"maximumTtl" yaml:"maximumTtl"`
	// The minimum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#minimum_ttl AwsLightsailDistribution#minimum_ttl}
	// Experimental.
	MinimumTtl *float64 `field:"optional" json:"minimumTtl" yaml:"minimumTtl"`
}

