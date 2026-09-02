package awsec2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSpotFleetRequestConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#iam_fleet_role TfSpotFleetRequest#iam_fleet_role}.
	// Experimental.
	IamFleetRole *string `field:"required" json:"iamFleetRole" yaml:"iamFleetRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#target_capacity TfSpotFleetRequest#target_capacity}.
	// Experimental.
	TargetCapacity *float64 `field:"required" json:"targetCapacity" yaml:"targetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#allocation_strategy TfSpotFleetRequest#allocation_strategy}.
	// Experimental.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#context TfSpotFleetRequest#context}.
	// Experimental.
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#excess_capacity_termination_policy TfSpotFleetRequest#excess_capacity_termination_policy}.
	// Experimental.
	ExcessCapacityTerminationPolicy *string `field:"optional" json:"excessCapacityTerminationPolicy" yaml:"excessCapacityTerminationPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#fleet_type TfSpotFleetRequest#fleet_type}.
	// Experimental.
	FleetType *string `field:"optional" json:"fleetType" yaml:"fleetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#id TfSpotFleetRequest#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#instance_interruption_behaviour TfSpotFleetRequest#instance_interruption_behaviour}.
	// Experimental.
	InstanceInterruptionBehaviour *string `field:"optional" json:"instanceInterruptionBehaviour" yaml:"instanceInterruptionBehaviour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#instance_pools_to_use_count TfSpotFleetRequest#instance_pools_to_use_count}.
	// Experimental.
	InstancePoolsToUseCount *float64 `field:"optional" json:"instancePoolsToUseCount" yaml:"instancePoolsToUseCount"`
	// launch_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#launch_specification TfSpotFleetRequest#launch_specification}
	// Experimental.
	LaunchSpecification interface{} `field:"optional" json:"launchSpecification" yaml:"launchSpecification"`
	// launch_template_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#launch_template_config TfSpotFleetRequest#launch_template_config}
	// Experimental.
	LaunchTemplateConfig interface{} `field:"optional" json:"launchTemplateConfig" yaml:"launchTemplateConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#load_balancers TfSpotFleetRequest#load_balancers}.
	// Experimental.
	LoadBalancers *[]*string `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#on_demand_allocation_strategy TfSpotFleetRequest#on_demand_allocation_strategy}.
	// Experimental.
	OnDemandAllocationStrategy *string `field:"optional" json:"onDemandAllocationStrategy" yaml:"onDemandAllocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#on_demand_max_total_price TfSpotFleetRequest#on_demand_max_total_price}.
	// Experimental.
	OnDemandMaxTotalPrice *string `field:"optional" json:"onDemandMaxTotalPrice" yaml:"onDemandMaxTotalPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#on_demand_target_capacity TfSpotFleetRequest#on_demand_target_capacity}.
	// Experimental.
	OnDemandTargetCapacity *float64 `field:"optional" json:"onDemandTargetCapacity" yaml:"onDemandTargetCapacity"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#region TfSpotFleetRequest#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#replace_unhealthy_instances TfSpotFleetRequest#replace_unhealthy_instances}.
	// Experimental.
	ReplaceUnhealthyInstances interface{} `field:"optional" json:"replaceUnhealthyInstances" yaml:"replaceUnhealthyInstances"`
	// spot_maintenance_strategies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#spot_maintenance_strategies TfSpotFleetRequest#spot_maintenance_strategies}
	// Experimental.
	SpotMaintenanceStrategies *TfSpotFleetRequest_SpotMaintenanceStrategiesProperty `field:"optional" json:"spotMaintenanceStrategies" yaml:"spotMaintenanceStrategies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#spot_price TfSpotFleetRequest#spot_price}.
	// Experimental.
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#tags TfSpotFleetRequest#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#tags_all TfSpotFleetRequest#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#target_capacity_unit_type TfSpotFleetRequest#target_capacity_unit_type}.
	// Experimental.
	TargetCapacityUnitType *string `field:"optional" json:"targetCapacityUnitType" yaml:"targetCapacityUnitType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#target_group_arns TfSpotFleetRequest#target_group_arns}.
	// Experimental.
	TargetGroupArns *[]*string `field:"optional" json:"targetGroupArns" yaml:"targetGroupArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#terminate_instances_on_delete TfSpotFleetRequest#terminate_instances_on_delete}.
	// Experimental.
	TerminateInstancesOnDelete *string `field:"optional" json:"terminateInstancesOnDelete" yaml:"terminateInstancesOnDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#terminate_instances_with_expiration TfSpotFleetRequest#terminate_instances_with_expiration}.
	// Experimental.
	TerminateInstancesWithExpiration interface{} `field:"optional" json:"terminateInstancesWithExpiration" yaml:"terminateInstancesWithExpiration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#timeouts TfSpotFleetRequest#timeouts}
	// Experimental.
	Timeouts *TfSpotFleetRequest_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#valid_from TfSpotFleetRequest#valid_from}.
	// Experimental.
	ValidFrom *string `field:"optional" json:"validFrom" yaml:"validFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#valid_until TfSpotFleetRequest#valid_until}.
	// Experimental.
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#wait_for_fulfillment TfSpotFleetRequest#wait_for_fulfillment}.
	// Experimental.
	WaitForFulfillment interface{} `field:"optional" json:"waitForFulfillment" yaml:"waitForFulfillment"`
}

