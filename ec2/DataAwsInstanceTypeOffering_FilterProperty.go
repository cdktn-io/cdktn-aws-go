package ec2


// Experimental.
type DataAwsInstanceTypeOffering_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_instance_type_offering#name DataAwsInstanceTypeOffering#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_instance_type_offering#values DataAwsInstanceTypeOffering#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

