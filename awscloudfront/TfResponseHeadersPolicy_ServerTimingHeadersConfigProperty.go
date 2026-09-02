package awscloudfront


// Experimental.
type TfResponseHeadersPolicy_ServerTimingHeadersConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#enabled TfResponseHeadersPolicy#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#sampling_rate TfResponseHeadersPolicy#sampling_rate}.
	// Experimental.
	SamplingRate *float64 `field:"required" json:"samplingRate" yaml:"samplingRate"`
}

