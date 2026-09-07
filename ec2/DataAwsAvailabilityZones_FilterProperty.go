package ec2


// Experimental.
type DataAwsAvailabilityZones_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/availability_zones#name DataAwsAvailabilityZones#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/availability_zones#values DataAwsAvailabilityZones#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

