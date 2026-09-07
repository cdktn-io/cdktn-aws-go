package cloudfront


// Experimental.
type AwsResponseHeadersPolicy_FrameOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#frame_option AwsResponseHeadersPolicy#frame_option}.
	// Experimental.
	FrameOption *string `field:"required" json:"frameOption" yaml:"frameOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#override AwsResponseHeadersPolicy#override}.
	// Experimental.
	Override interface{} `field:"required" json:"override" yaml:"override"`
}

