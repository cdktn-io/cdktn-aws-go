package awsec2


// Experimental.
type DataTfCapacityBlockReservation_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_capacity_block_reservation#name DataTfCapacityBlockReservation#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_capacity_block_reservation#values DataTfCapacityBlockReservation#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

