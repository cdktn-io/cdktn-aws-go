package awscloudfront


// Experimental.
type TfResponseHeadersPolicy_StrictTransportSecurityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#access_control_max_age_sec TfResponseHeadersPolicy#access_control_max_age_sec}.
	// Experimental.
	AccessControlMaxAgeSec *float64 `field:"required" json:"accessControlMaxAgeSec" yaml:"accessControlMaxAgeSec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#override TfResponseHeadersPolicy#override}.
	// Experimental.
	Override interface{} `field:"required" json:"override" yaml:"override"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#include_subdomains TfResponseHeadersPolicy#include_subdomains}.
	// Experimental.
	IncludeSubdomains interface{} `field:"optional" json:"includeSubdomains" yaml:"includeSubdomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#preload TfResponseHeadersPolicy#preload}.
	// Experimental.
	Preload interface{} `field:"optional" json:"preload" yaml:"preload"`
}

