package awsemr


// Experimental.
type TfCluster_MasterInstanceFleetLaunchSpecificationsProperty struct {
	// on_demand_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#on_demand_specification TfCluster#on_demand_specification}
	// Experimental.
	OnDemandSpecification interface{} `field:"optional" json:"onDemandSpecification" yaml:"onDemandSpecification"`
	// spot_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#spot_specification TfCluster#spot_specification}
	// Experimental.
	SpotSpecification interface{} `field:"optional" json:"spotSpecification" yaml:"spotSpecification"`
}

