package cloudfront


// Experimental.
type AwsResponseHeadersPolicy_CorsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_credentials AwsResponseHeadersPolicy#access_control_allow_credentials}.
	// Experimental.
	AccessControlAllowCredentials interface{} `field:"required" json:"accessControlAllowCredentials" yaml:"accessControlAllowCredentials"`
	// access_control_allow_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_headers AwsResponseHeadersPolicy#access_control_allow_headers}
	// Experimental.
	AccessControlAllowHeaders *AwsResponseHeadersPolicy_AccessControlAllowHeadersProperty `field:"required" json:"accessControlAllowHeaders" yaml:"accessControlAllowHeaders"`
	// access_control_allow_methods block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_methods AwsResponseHeadersPolicy#access_control_allow_methods}
	// Experimental.
	AccessControlAllowMethods *AwsResponseHeadersPolicy_AccessControlAllowMethodsProperty `field:"required" json:"accessControlAllowMethods" yaml:"accessControlAllowMethods"`
	// access_control_allow_origins block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_origins AwsResponseHeadersPolicy#access_control_allow_origins}
	// Experimental.
	AccessControlAllowOrigins *AwsResponseHeadersPolicy_AccessControlAllowOriginsProperty `field:"required" json:"accessControlAllowOrigins" yaml:"accessControlAllowOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#origin_override AwsResponseHeadersPolicy#origin_override}.
	// Experimental.
	OriginOverride interface{} `field:"required" json:"originOverride" yaml:"originOverride"`
	// access_control_expose_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_expose_headers AwsResponseHeadersPolicy#access_control_expose_headers}
	// Experimental.
	AccessControlExposeHeaders *AwsResponseHeadersPolicy_AccessControlExposeHeadersProperty `field:"optional" json:"accessControlExposeHeaders" yaml:"accessControlExposeHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_max_age_sec AwsResponseHeadersPolicy#access_control_max_age_sec}.
	// Experimental.
	AccessControlMaxAgeSec *float64 `field:"optional" json:"accessControlMaxAgeSec" yaml:"accessControlMaxAgeSec"`
}

