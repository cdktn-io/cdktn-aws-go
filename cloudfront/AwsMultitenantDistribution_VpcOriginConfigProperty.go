package cloudfront


// Experimental.
type AwsMultitenantDistribution_VpcOriginConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#vpc_origin_id AwsMultitenantDistribution#vpc_origin_id}.
	// Experimental.
	VpcOriginId *string `field:"required" json:"vpcOriginId" yaml:"vpcOriginId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin_keepalive_timeout AwsMultitenantDistribution#origin_keepalive_timeout}.
	// Experimental.
	OriginKeepaliveTimeout *float64 `field:"optional" json:"originKeepaliveTimeout" yaml:"originKeepaliveTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin_read_timeout AwsMultitenantDistribution#origin_read_timeout}.
	// Experimental.
	OriginReadTimeout *float64 `field:"optional" json:"originReadTimeout" yaml:"originReadTimeout"`
}

