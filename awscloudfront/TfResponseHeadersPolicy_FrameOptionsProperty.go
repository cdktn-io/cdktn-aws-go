package awscloudfront


// Experimental.
type TfResponseHeadersPolicy_FrameOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#frame_option TfResponseHeadersPolicy#frame_option}.
	// Experimental.
	FrameOption *string `field:"required" json:"frameOption" yaml:"frameOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#override TfResponseHeadersPolicy#override}.
	// Experimental.
	Override interface{} `field:"required" json:"override" yaml:"override"`
}

