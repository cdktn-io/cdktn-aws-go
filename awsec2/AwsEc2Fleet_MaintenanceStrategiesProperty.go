package awsec2


// Experimental.
type AwsEc2Fleet_MaintenanceStrategiesProperty struct {
	// capacity_rebalance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#capacity_rebalance AwsEc2Fleet#capacity_rebalance}
	// Experimental.
	CapacityRebalance *AwsEc2Fleet_CapacityRebalanceProperty `field:"optional" json:"capacityRebalance" yaml:"capacityRebalance"`
}

