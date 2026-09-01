package awsroute53


// Experimental.
type AwsRoute53Zone_VpcProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_zone#vpc_id AwsRoute53Zone#vpc_id}.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_zone#vpc_region AwsRoute53Zone#vpc_region}.
	// Experimental.
	VpcRegion *string `field:"optional" json:"vpcRegion" yaml:"vpcRegion"`
}

