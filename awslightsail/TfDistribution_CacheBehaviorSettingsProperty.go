package awslightsail


// Experimental.
type TfDistribution_CacheBehaviorSettingsProperty struct {
	// The HTTP methods that are processed and forwarded to the distribution's origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#allowed_http_methods TfDistribution#allowed_http_methods}
	// Experimental.
	AllowedHttpMethods *string `field:"optional" json:"allowedHttpMethods" yaml:"allowedHttpMethods"`
	// The HTTP method responses that are cached by your distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#cached_http_methods TfDistribution#cached_http_methods}
	// Experimental.
	CachedHttpMethods *string `field:"optional" json:"cachedHttpMethods" yaml:"cachedHttpMethods"`
	// The default amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the content has been updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#default_ttl TfDistribution#default_ttl}
	// Experimental.
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// forwarded_cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#forwarded_cookies TfDistribution#forwarded_cookies}
	// Experimental.
	ForwardedCookies *TfDistribution_ForwardedCookiesProperty `field:"optional" json:"forwardedCookies" yaml:"forwardedCookies"`
	// forwarded_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#forwarded_headers TfDistribution#forwarded_headers}
	// Experimental.
	ForwardedHeaders *TfDistribution_ForwardedHeadersProperty `field:"optional" json:"forwardedHeaders" yaml:"forwardedHeaders"`
	// forwarded_query_strings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#forwarded_query_strings TfDistribution#forwarded_query_strings}
	// Experimental.
	ForwardedQueryStrings *TfDistribution_ForwardedQueryStringsProperty `field:"optional" json:"forwardedQueryStrings" yaml:"forwardedQueryStrings"`
	// The maximum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#maximum_ttl TfDistribution#maximum_ttl}
	// Experimental.
	MaximumTtl *float64 `field:"optional" json:"maximumTtl" yaml:"maximumTtl"`
	// The minimum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#minimum_ttl TfDistribution#minimum_ttl}
	// Experimental.
	MinimumTtl *float64 `field:"optional" json:"minimumTtl" yaml:"minimumTtl"`
}

