package awscloudfront


// Experimental.
type TfMultitenantDistribution_CustomHeaderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#header_name TfMultitenantDistribution#header_name}.
	// Experimental.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#header_value TfMultitenantDistribution#header_value}.
	// Experimental.
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
}

