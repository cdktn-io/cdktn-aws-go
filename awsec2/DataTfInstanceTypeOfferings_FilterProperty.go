package awsec2


// Experimental.
type DataTfInstanceTypeOfferings_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_instance_type_offerings#name DataTfInstanceTypeOfferings#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_instance_type_offerings#values DataTfInstanceTypeOfferings#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

