package cloudfront


// Experimental.
type AwsResponseHeadersPolicy_XssProtectionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#override AwsResponseHeadersPolicy#override}.
	// Experimental.
	Override interface{} `field:"required" json:"override" yaml:"override"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#protection AwsResponseHeadersPolicy#protection}.
	// Experimental.
	Protection interface{} `field:"required" json:"protection" yaml:"protection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#mode_block AwsResponseHeadersPolicy#mode_block}.
	// Experimental.
	ModeBlock interface{} `field:"optional" json:"modeBlock" yaml:"modeBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#report_uri AwsResponseHeadersPolicy#report_uri}.
	// Experimental.
	ReportUri *string `field:"optional" json:"reportUri" yaml:"reportUri"`
}

