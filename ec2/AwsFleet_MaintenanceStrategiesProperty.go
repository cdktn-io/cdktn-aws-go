package ec2


// Experimental.
type AwsFleet_MaintenanceStrategiesProperty struct {
	// capacity_rebalance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#capacity_rebalance AwsFleet#capacity_rebalance}
	// Experimental.
	CapacityRebalance *AwsFleet_CapacityRebalanceProperty `field:"optional" json:"capacityRebalance" yaml:"capacityRebalance"`
}

