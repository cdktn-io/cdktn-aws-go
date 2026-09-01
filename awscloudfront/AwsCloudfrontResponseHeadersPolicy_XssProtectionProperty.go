package awscloudfront


// Experimental.
type AwsCloudfrontResponseHeadersPolicy_XssProtectionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#override AwsCloudfrontResponseHeadersPolicy#override}.
	// Experimental.
	Override interface{} `field:"required" json:"override" yaml:"override"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#protection AwsCloudfrontResponseHeadersPolicy#protection}.
	// Experimental.
	Protection interface{} `field:"required" json:"protection" yaml:"protection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#mode_block AwsCloudfrontResponseHeadersPolicy#mode_block}.
	// Experimental.
	ModeBlock interface{} `field:"optional" json:"modeBlock" yaml:"modeBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#report_uri AwsCloudfrontResponseHeadersPolicy#report_uri}.
	// Experimental.
	ReportUri *string `field:"optional" json:"reportUri" yaml:"reportUri"`
}

