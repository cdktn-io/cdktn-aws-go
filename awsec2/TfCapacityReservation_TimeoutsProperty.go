package awsec2


// Experimental.
type TfCapacityReservation_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_capacity_reservation#create TfCapacityReservation#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_capacity_reservation#delete TfCapacityReservation#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_capacity_reservation#update TfCapacityReservation#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

