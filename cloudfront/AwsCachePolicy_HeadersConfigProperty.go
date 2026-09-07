package cloudfront


// Experimental.
type AwsCachePolicy_HeadersConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#header_behavior AwsCachePolicy#header_behavior}.
	// Experimental.
	HeaderBehavior *string `field:"optional" json:"headerBehavior" yaml:"headerBehavior"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#headers AwsCachePolicy#headers}
	// Experimental.
	Headers *AwsCachePolicy_HeadersProperty `field:"optional" json:"headers" yaml:"headers"`
}

