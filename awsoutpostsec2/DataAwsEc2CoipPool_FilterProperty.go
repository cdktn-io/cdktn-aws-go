package awsoutpostsec2


// Experimental.
type DataAwsEc2CoipPool_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_coip_pool#name DataAwsEc2CoipPool#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_coip_pool#values DataAwsEc2CoipPool#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

