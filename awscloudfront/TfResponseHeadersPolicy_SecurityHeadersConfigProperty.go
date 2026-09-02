package awscloudfront


// Experimental.
type TfResponseHeadersPolicy_SecurityHeadersConfigProperty struct {
	// content_security_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#content_security_policy TfResponseHeadersPolicy#content_security_policy}
	// Experimental.
	ContentSecurityPolicy *TfResponseHeadersPolicy_ContentSecurityPolicyProperty `field:"optional" json:"contentSecurityPolicy" yaml:"contentSecurityPolicy"`
	// content_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#content_type_options TfResponseHeadersPolicy#content_type_options}
	// Experimental.
	ContentTypeOptions *TfResponseHeadersPolicy_ContentTypeOptionsProperty `field:"optional" json:"contentTypeOptions" yaml:"contentTypeOptions"`
	// frame_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#frame_options TfResponseHeadersPolicy#frame_options}
	// Experimental.
	FrameOptions *TfResponseHeadersPolicy_FrameOptionsProperty `field:"optional" json:"frameOptions" yaml:"frameOptions"`
	// referrer_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#referrer_policy TfResponseHeadersPolicy#referrer_policy}
	// Experimental.
	ReferrerPolicy *TfResponseHeadersPolicy_ReferrerPolicyProperty `field:"optional" json:"referrerPolicy" yaml:"referrerPolicy"`
	// strict_transport_security block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#strict_transport_security TfResponseHeadersPolicy#strict_transport_security}
	// Experimental.
	StrictTransportSecurity *TfResponseHeadersPolicy_StrictTransportSecurityProperty `field:"optional" json:"strictTransportSecurity" yaml:"strictTransportSecurity"`
	// xss_protection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#xss_protection TfResponseHeadersPolicy#xss_protection}
	// Experimental.
	XssProtection *TfResponseHeadersPolicy_XssProtectionProperty `field:"optional" json:"xssProtection" yaml:"xssProtection"`
}

