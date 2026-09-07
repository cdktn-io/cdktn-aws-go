package cloudfront


// Experimental.
type AwsDistribution_VpcOriginConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#vpc_origin_id AwsDistribution#vpc_origin_id}.
	// Experimental.
	VpcOriginId *string `field:"required" json:"vpcOriginId" yaml:"vpcOriginId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#origin_keepalive_timeout AwsDistribution#origin_keepalive_timeout}.
	// Experimental.
	OriginKeepaliveTimeout *float64 `field:"optional" json:"originKeepaliveTimeout" yaml:"originKeepaliveTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#origin_read_timeout AwsDistribution#origin_read_timeout}.
	// Experimental.
	OriginReadTimeout *float64 `field:"optional" json:"originReadTimeout" yaml:"originReadTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#owner_account_id AwsDistribution#owner_account_id}.
	// Experimental.
	OwnerAccountId *string `field:"optional" json:"ownerAccountId" yaml:"ownerAccountId"`
}

