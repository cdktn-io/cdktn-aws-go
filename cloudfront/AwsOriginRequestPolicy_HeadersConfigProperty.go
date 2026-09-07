package cloudfront


// Experimental.
type AwsOriginRequestPolicy_HeadersConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_origin_request_policy#header_behavior AwsOriginRequestPolicy#header_behavior}.
	// Experimental.
	HeaderBehavior *string `field:"optional" json:"headerBehavior" yaml:"headerBehavior"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_origin_request_policy#headers AwsOriginRequestPolicy#headers}
	// Experimental.
	Headers *AwsOriginRequestPolicy_HeadersProperty `field:"optional" json:"headers" yaml:"headers"`
}

