package awsemr


// Experimental.
type TfCluster_CoreInstanceFleetProperty struct {
	// instance_type_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#instance_type_configs TfCluster#instance_type_configs}
	// Experimental.
	InstanceTypeConfigs interface{} `field:"optional" json:"instanceTypeConfigs" yaml:"instanceTypeConfigs"`
	// launch_specifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#launch_specifications TfCluster#launch_specifications}
	// Experimental.
	LaunchSpecifications *TfCluster_CoreInstanceFleetLaunchSpecificationsProperty `field:"optional" json:"launchSpecifications" yaml:"launchSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#name TfCluster#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#target_on_demand_capacity TfCluster#target_on_demand_capacity}.
	// Experimental.
	TargetOnDemandCapacity *float64 `field:"optional" json:"targetOnDemandCapacity" yaml:"targetOnDemandCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#target_spot_capacity TfCluster#target_spot_capacity}.
	// Experimental.
	TargetSpotCapacity *float64 `field:"optional" json:"targetSpotCapacity" yaml:"targetSpotCapacity"`
}

