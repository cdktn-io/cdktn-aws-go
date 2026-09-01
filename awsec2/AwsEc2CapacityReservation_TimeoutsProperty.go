package awsec2


// Experimental.
type AwsEc2CapacityReservation_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_capacity_reservation#create AwsEc2CapacityReservation#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_capacity_reservation#delete AwsEc2CapacityReservation#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_capacity_reservation#update AwsEc2CapacityReservation#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

