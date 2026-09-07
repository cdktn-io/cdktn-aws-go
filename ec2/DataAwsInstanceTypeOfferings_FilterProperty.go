package ec2


// Experimental.
type DataAwsInstanceTypeOfferings_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_instance_type_offerings#name DataAwsInstanceTypeOfferings#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_instance_type_offerings#values DataAwsInstanceTypeOfferings#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

