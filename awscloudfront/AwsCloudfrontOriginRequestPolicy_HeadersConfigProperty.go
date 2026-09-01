package awscloudfront


// Experimental.
type AwsCloudfrontOriginRequestPolicy_HeadersConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_origin_request_policy#header_behavior AwsCloudfrontOriginRequestPolicy#header_behavior}.
	// Experimental.
	HeaderBehavior *string `field:"optional" json:"headerBehavior" yaml:"headerBehavior"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_origin_request_policy#headers AwsCloudfrontOriginRequestPolicy#headers}
	// Experimental.
	Headers *AwsCloudfrontOriginRequestPolicy_HeadersProperty `field:"optional" json:"headers" yaml:"headers"`
}

