package awscloudfront


// Experimental.
type AwsCloudfrontCachePolicy_HeadersConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#header_behavior AwsCloudfrontCachePolicy#header_behavior}.
	// Experimental.
	HeaderBehavior *string `field:"optional" json:"headerBehavior" yaml:"headerBehavior"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#headers AwsCloudfrontCachePolicy#headers}
	// Experimental.
	Headers *AwsCloudfrontCachePolicy_HeadersProperty `field:"optional" json:"headers" yaml:"headers"`
}

