package ec2


// Experimental.
type DataAwsInstanceTypes_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_instance_types#name DataAwsInstanceTypes#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_instance_types#values DataAwsInstanceTypes#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

