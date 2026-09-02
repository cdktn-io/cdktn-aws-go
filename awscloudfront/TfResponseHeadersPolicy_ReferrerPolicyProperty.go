package awscloudfront


// Experimental.
type TfResponseHeadersPolicy_ReferrerPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#override TfResponseHeadersPolicy#override}.
	// Experimental.
	Override interface{} `field:"required" json:"override" yaml:"override"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#referrer_policy TfResponseHeadersPolicy#referrer_policy}.
	// Experimental.
	ReferrerPolicy *string `field:"required" json:"referrerPolicy" yaml:"referrerPolicy"`
}

