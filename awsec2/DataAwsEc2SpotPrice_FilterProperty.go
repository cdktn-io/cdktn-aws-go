package awsec2


// Experimental.
type DataAwsEc2SpotPrice_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_spot_price#name DataAwsEc2SpotPrice#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_spot_price#values DataAwsEc2SpotPrice#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

