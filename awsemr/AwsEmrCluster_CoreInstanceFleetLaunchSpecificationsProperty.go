package awsemr


// Experimental.
type AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsProperty struct {
	// on_demand_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#on_demand_specification AwsEmrCluster#on_demand_specification}
	// Experimental.
	OnDemandSpecification interface{} `field:"optional" json:"onDemandSpecification" yaml:"onDemandSpecification"`
	// spot_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#spot_specification AwsEmrCluster#spot_specification}
	// Experimental.
	SpotSpecification interface{} `field:"optional" json:"spotSpecification" yaml:"spotSpecification"`
}

