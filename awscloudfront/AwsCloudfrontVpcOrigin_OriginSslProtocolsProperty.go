package awscloudfront


// Experimental.
type AwsCloudfrontVpcOrigin_OriginSslProtocolsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_vpc_origin#items AwsCloudfrontVpcOrigin#items}.
	// Experimental.
	Items *[]*string `field:"required" json:"items" yaml:"items"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_vpc_origin#quantity AwsCloudfrontVpcOrigin#quantity}.
	// Experimental.
	Quantity *float64 `field:"required" json:"quantity" yaml:"quantity"`
}

