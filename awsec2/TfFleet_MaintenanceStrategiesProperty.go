package awsec2


// Experimental.
type TfFleet_MaintenanceStrategiesProperty struct {
	// capacity_rebalance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#capacity_rebalance TfFleet#capacity_rebalance}
	// Experimental.
	CapacityRebalance *TfFleet_CapacityRebalanceProperty `field:"optional" json:"capacityRebalance" yaml:"capacityRebalance"`
}

