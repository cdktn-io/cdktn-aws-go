package awscloudfront


// Experimental.
type TfResponseHeadersPolicy_ContentSecurityPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#content_security_policy TfResponseHeadersPolicy#content_security_policy}.
	// Experimental.
	ContentSecurityPolicy *string `field:"required" json:"contentSecurityPolicy" yaml:"contentSecurityPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#override TfResponseHeadersPolicy#override}.
	// Experimental.
	Override interface{} `field:"required" json:"override" yaml:"override"`
}

