package awscloudfront


// Experimental.
type AwsCloudfrontResponseHeadersPolicy_CorsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_credentials AwsCloudfrontResponseHeadersPolicy#access_control_allow_credentials}.
	// Experimental.
	AccessControlAllowCredentials interface{} `field:"required" json:"accessControlAllowCredentials" yaml:"accessControlAllowCredentials"`
	// access_control_allow_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_headers AwsCloudfrontResponseHeadersPolicy#access_control_allow_headers}
	// Experimental.
	AccessControlAllowHeaders *AwsCloudfrontResponseHeadersPolicy_AccessControlAllowHeadersProperty `field:"required" json:"accessControlAllowHeaders" yaml:"accessControlAllowHeaders"`
	// access_control_allow_methods block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_methods AwsCloudfrontResponseHeadersPolicy#access_control_allow_methods}
	// Experimental.
	AccessControlAllowMethods *AwsCloudfrontResponseHeadersPolicy_AccessControlAllowMethodsProperty `field:"required" json:"accessControlAllowMethods" yaml:"accessControlAllowMethods"`
	// access_control_allow_origins block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_origins AwsCloudfrontResponseHeadersPolicy#access_control_allow_origins}
	// Experimental.
	AccessControlAllowOrigins *AwsCloudfrontResponseHeadersPolicy_AccessControlAllowOriginsProperty `field:"required" json:"accessControlAllowOrigins" yaml:"accessControlAllowOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#origin_override AwsCloudfrontResponseHeadersPolicy#origin_override}.
	// Experimental.
	OriginOverride interface{} `field:"required" json:"originOverride" yaml:"originOverride"`
	// access_control_expose_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_expose_headers AwsCloudfrontResponseHeadersPolicy#access_control_expose_headers}
	// Experimental.
	AccessControlExposeHeaders *AwsCloudfrontResponseHeadersPolicy_AccessControlExposeHeadersProperty `field:"optional" json:"accessControlExposeHeaders" yaml:"accessControlExposeHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_max_age_sec AwsCloudfrontResponseHeadersPolicy#access_control_max_age_sec}.
	// Experimental.
	AccessControlMaxAgeSec *float64 `field:"optional" json:"accessControlMaxAgeSec" yaml:"accessControlMaxAgeSec"`
}

