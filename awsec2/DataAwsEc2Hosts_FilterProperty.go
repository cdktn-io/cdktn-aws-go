package awsec2


// Experimental.
type DataAwsEc2Hosts_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_hosts#name DataAwsEc2Hosts#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_hosts#values DataAwsEc2Hosts#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

