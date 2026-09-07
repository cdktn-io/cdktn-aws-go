package ec2


// Experimental.
type DataAwsInstances_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/instances#name DataAwsInstances#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/instances#values DataAwsInstances#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

