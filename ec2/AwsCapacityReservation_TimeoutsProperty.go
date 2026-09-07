package ec2


// Experimental.
type AwsCapacityReservation_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_capacity_reservation#create AwsCapacityReservation#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_capacity_reservation#delete AwsCapacityReservation#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_capacity_reservation#update AwsCapacityReservation#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

