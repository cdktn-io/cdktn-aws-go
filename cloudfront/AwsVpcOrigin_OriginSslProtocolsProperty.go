package cloudfront


// Experimental.
type AwsVpcOrigin_OriginSslProtocolsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_vpc_origin#items AwsVpcOrigin#items}.
	// Experimental.
	Items *[]*string `field:"required" json:"items" yaml:"items"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_vpc_origin#quantity AwsVpcOrigin#quantity}.
	// Experimental.
	Quantity *float64 `field:"required" json:"quantity" yaml:"quantity"`
}

