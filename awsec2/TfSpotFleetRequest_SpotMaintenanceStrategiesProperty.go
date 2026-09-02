package awsec2


// Experimental.
type TfSpotFleetRequest_SpotMaintenanceStrategiesProperty struct {
	// capacity_rebalance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#capacity_rebalance TfSpotFleetRequest#capacity_rebalance}
	// Experimental.
	CapacityRebalance *TfSpotFleetRequest_CapacityRebalanceProperty `field:"optional" json:"capacityRebalance" yaml:"capacityRebalance"`
}

