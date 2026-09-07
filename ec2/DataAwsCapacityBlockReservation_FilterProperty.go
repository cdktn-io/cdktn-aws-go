package ec2


// Experimental.
type DataAwsCapacityBlockReservation_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_capacity_block_reservation#name DataAwsCapacityBlockReservation#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_capacity_block_reservation#values DataAwsCapacityBlockReservation#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

