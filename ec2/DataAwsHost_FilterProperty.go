package ec2


// Experimental.
type DataAwsHost_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_host#name DataAwsHost#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_host#values DataAwsHost#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

