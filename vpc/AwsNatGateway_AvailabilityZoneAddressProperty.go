package vpc


// Experimental.
type AwsNatGateway_AvailabilityZoneAddressProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#allocation_ids AwsNatGateway#allocation_ids}.
	// Experimental.
	AllocationIds *[]*string `field:"optional" json:"allocationIds" yaml:"allocationIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#availability_zone AwsNatGateway#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#availability_zone_id AwsNatGateway#availability_zone_id}.
	// Experimental.
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
}

