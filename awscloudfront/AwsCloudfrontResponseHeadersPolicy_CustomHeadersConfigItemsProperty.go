package awscloudfront


// Experimental.
type AwsCloudfrontResponseHeadersPolicy_CustomHeadersConfigItemsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#header AwsCloudfrontResponseHeadersPolicy#header}.
	// Experimental.
	Header *string `field:"required" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#override AwsCloudfrontResponseHeadersPolicy#override}.
	// Experimental.
	Override interface{} `field:"required" json:"override" yaml:"override"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#value AwsCloudfrontResponseHeadersPolicy#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

