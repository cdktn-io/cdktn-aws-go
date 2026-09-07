package emr


// Experimental.
type AwsCluster_MasterInstanceFleetProperty struct {
	// instance_type_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#instance_type_configs AwsCluster#instance_type_configs}
	// Experimental.
	InstanceTypeConfigs interface{} `field:"optional" json:"instanceTypeConfigs" yaml:"instanceTypeConfigs"`
	// launch_specifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#launch_specifications AwsCluster#launch_specifications}
	// Experimental.
	LaunchSpecifications *AwsCluster_MasterInstanceFleetLaunchSpecificationsProperty `field:"optional" json:"launchSpecifications" yaml:"launchSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#name AwsCluster#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#target_on_demand_capacity AwsCluster#target_on_demand_capacity}.
	// Experimental.
	TargetOnDemandCapacity *float64 `field:"optional" json:"targetOnDemandCapacity" yaml:"targetOnDemandCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#target_spot_capacity AwsCluster#target_spot_capacity}.
	// Experimental.
	TargetSpotCapacity *float64 `field:"optional" json:"targetSpotCapacity" yaml:"targetSpotCapacity"`
}

