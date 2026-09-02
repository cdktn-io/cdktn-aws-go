package awsroute53


// Experimental.
type TfZone_VpcProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_zone#vpc_id TfZone#vpc_id}.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_zone#vpc_region TfZone#vpc_region}.
	// Experimental.
	VpcRegion *string `field:"optional" json:"vpcRegion" yaml:"vpcRegion"`
}

