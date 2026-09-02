package awscloudfront


// Experimental.
type TfResponseHeadersPolicy_CorsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_credentials TfResponseHeadersPolicy#access_control_allow_credentials}.
	// Experimental.
	AccessControlAllowCredentials interface{} `field:"required" json:"accessControlAllowCredentials" yaml:"accessControlAllowCredentials"`
	// access_control_allow_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_headers TfResponseHeadersPolicy#access_control_allow_headers}
	// Experimental.
	AccessControlAllowHeaders *TfResponseHeadersPolicy_AccessControlAllowHeadersProperty `field:"required" json:"accessControlAllowHeaders" yaml:"accessControlAllowHeaders"`
	// access_control_allow_methods block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_methods TfResponseHeadersPolicy#access_control_allow_methods}
	// Experimental.
	AccessControlAllowMethods *TfResponseHeadersPolicy_AccessControlAllowMethodsProperty `field:"required" json:"accessControlAllowMethods" yaml:"accessControlAllowMethods"`
	// access_control_allow_origins block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_origins TfResponseHeadersPolicy#access_control_allow_origins}
	// Experimental.
	AccessControlAllowOrigins *TfResponseHeadersPolicy_AccessControlAllowOriginsProperty `field:"required" json:"accessControlAllowOrigins" yaml:"accessControlAllowOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#origin_override TfResponseHeadersPolicy#origin_override}.
	// Experimental.
	OriginOverride interface{} `field:"required" json:"originOverride" yaml:"originOverride"`
	// access_control_expose_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_expose_headers TfResponseHeadersPolicy#access_control_expose_headers}
	// Experimental.
	AccessControlExposeHeaders *TfResponseHeadersPolicy_AccessControlExposeHeadersProperty `field:"optional" json:"accessControlExposeHeaders" yaml:"accessControlExposeHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_max_age_sec TfResponseHeadersPolicy#access_control_max_age_sec}.
	// Experimental.
	AccessControlMaxAgeSec *float64 `field:"optional" json:"accessControlMaxAgeSec" yaml:"accessControlMaxAgeSec"`
}

