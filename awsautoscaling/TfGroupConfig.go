package awsautoscaling

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#max_size TfGroup#max_size}.
	// Experimental.
	MaxSize *float64 `field:"required" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#min_size TfGroup#min_size}.
	// Experimental.
	MinSize *float64 `field:"required" json:"minSize" yaml:"minSize"`
	// availability_zone_distribution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#availability_zone_distribution TfGroup#availability_zone_distribution}
	// Experimental.
	AvailabilityZoneDistribution *TfGroup_AvailabilityZoneDistributionProperty `field:"optional" json:"availabilityZoneDistribution" yaml:"availabilityZoneDistribution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#availability_zones TfGroup#availability_zones}.
	// Experimental.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#capacity_rebalance TfGroup#capacity_rebalance}.
	// Experimental.
	CapacityRebalance interface{} `field:"optional" json:"capacityRebalance" yaml:"capacityRebalance"`
	// capacity_reservation_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#capacity_reservation_specification TfGroup#capacity_reservation_specification}
	// Experimental.
	CapacityReservationSpecification *TfGroup_CapacityReservationSpecificationProperty `field:"optional" json:"capacityReservationSpecification" yaml:"capacityReservationSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#context TfGroup#context}.
	// Experimental.
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#default_cooldown TfGroup#default_cooldown}.
	// Experimental.
	DefaultCooldown *float64 `field:"optional" json:"defaultCooldown" yaml:"defaultCooldown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#default_instance_warmup TfGroup#default_instance_warmup}.
	// Experimental.
	DefaultInstanceWarmup *float64 `field:"optional" json:"defaultInstanceWarmup" yaml:"defaultInstanceWarmup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#desired_capacity TfGroup#desired_capacity}.
	// Experimental.
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#desired_capacity_type TfGroup#desired_capacity_type}.
	// Experimental.
	DesiredCapacityType *string `field:"optional" json:"desiredCapacityType" yaml:"desiredCapacityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#enabled_metrics TfGroup#enabled_metrics}.
	// Experimental.
	EnabledMetrics *[]*string `field:"optional" json:"enabledMetrics" yaml:"enabledMetrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#force_delete TfGroup#force_delete}.
	// Experimental.
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#force_delete_warm_pool TfGroup#force_delete_warm_pool}.
	// Experimental.
	ForceDeleteWarmPool interface{} `field:"optional" json:"forceDeleteWarmPool" yaml:"forceDeleteWarmPool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#health_check_grace_period TfGroup#health_check_grace_period}.
	// Experimental.
	HealthCheckGracePeriod *float64 `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#health_check_type TfGroup#health_check_type}.
	// Experimental.
	HealthCheckType *string `field:"optional" json:"healthCheckType" yaml:"healthCheckType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#id TfGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#ignore_failed_scaling_activities TfGroup#ignore_failed_scaling_activities}.
	// Experimental.
	IgnoreFailedScalingActivities interface{} `field:"optional" json:"ignoreFailedScalingActivities" yaml:"ignoreFailedScalingActivities"`
	// initial_lifecycle_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#initial_lifecycle_hook TfGroup#initial_lifecycle_hook}
	// Experimental.
	InitialLifecycleHook interface{} `field:"optional" json:"initialLifecycleHook" yaml:"initialLifecycleHook"`
	// instance_lifecycle_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#instance_lifecycle_policy TfGroup#instance_lifecycle_policy}
	// Experimental.
	InstanceLifecyclePolicy *TfGroup_InstanceLifecyclePolicyProperty `field:"optional" json:"instanceLifecyclePolicy" yaml:"instanceLifecyclePolicy"`
	// instance_maintenance_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#instance_maintenance_policy TfGroup#instance_maintenance_policy}
	// Experimental.
	InstanceMaintenancePolicy *TfGroup_InstanceMaintenancePolicyProperty `field:"optional" json:"instanceMaintenancePolicy" yaml:"instanceMaintenancePolicy"`
	// instance_refresh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#instance_refresh TfGroup#instance_refresh}
	// Experimental.
	InstanceRefresh *TfGroup_InstanceRefreshProperty `field:"optional" json:"instanceRefresh" yaml:"instanceRefresh"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#launch_configuration TfGroup#launch_configuration}.
	// Experimental.
	LaunchConfiguration *string `field:"optional" json:"launchConfiguration" yaml:"launchConfiguration"`
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#launch_template TfGroup#launch_template}
	// Experimental.
	LaunchTemplate *TfGroup_LaunchTemplateProperty `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#load_balancers TfGroup#load_balancers}.
	// Experimental.
	LoadBalancers *[]*string `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#max_instance_lifetime TfGroup#max_instance_lifetime}.
	// Experimental.
	MaxInstanceLifetime *float64 `field:"optional" json:"maxInstanceLifetime" yaml:"maxInstanceLifetime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#metrics_granularity TfGroup#metrics_granularity}.
	// Experimental.
	MetricsGranularity *string `field:"optional" json:"metricsGranularity" yaml:"metricsGranularity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#min_elb_capacity TfGroup#min_elb_capacity}.
	// Experimental.
	MinElbCapacity *float64 `field:"optional" json:"minElbCapacity" yaml:"minElbCapacity"`
	// mixed_instances_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#mixed_instances_policy TfGroup#mixed_instances_policy}
	// Experimental.
	MixedInstancesPolicy *TfGroup_MixedInstancesPolicyProperty `field:"optional" json:"mixedInstancesPolicy" yaml:"mixedInstancesPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#name TfGroup#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#name_prefix TfGroup#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#placement_group TfGroup#placement_group}.
	// Experimental.
	PlacementGroup *string `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#protect_from_scale_in TfGroup#protect_from_scale_in}.
	// Experimental.
	ProtectFromScaleIn interface{} `field:"optional" json:"protectFromScaleIn" yaml:"protectFromScaleIn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#region TfGroup#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#service_linked_role_arn TfGroup#service_linked_role_arn}.
	// Experimental.
	ServiceLinkedRoleArn *string `field:"optional" json:"serviceLinkedRoleArn" yaml:"serviceLinkedRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#suspended_processes TfGroup#suspended_processes}.
	// Experimental.
	SuspendedProcesses *[]*string `field:"optional" json:"suspendedProcesses" yaml:"suspendedProcesses"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#tag TfGroup#tag}
	// Experimental.
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#target_group_arns TfGroup#target_group_arns}.
	// Experimental.
	TargetGroupArns *[]*string `field:"optional" json:"targetGroupArns" yaml:"targetGroupArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#termination_policies TfGroup#termination_policies}.
	// Experimental.
	TerminationPolicies *[]*string `field:"optional" json:"terminationPolicies" yaml:"terminationPolicies"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#timeouts TfGroup#timeouts}
	// Experimental.
	Timeouts *TfGroup_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// traffic_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#traffic_source TfGroup#traffic_source}
	// Experimental.
	TrafficSource interface{} `field:"optional" json:"trafficSource" yaml:"trafficSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#vpc_zone_identifier TfGroup#vpc_zone_identifier}.
	// Experimental.
	VpcZoneIdentifier *[]*string `field:"optional" json:"vpcZoneIdentifier" yaml:"vpcZoneIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#wait_for_capacity_timeout TfGroup#wait_for_capacity_timeout}.
	// Experimental.
	WaitForCapacityTimeout *string `field:"optional" json:"waitForCapacityTimeout" yaml:"waitForCapacityTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#wait_for_elb_capacity TfGroup#wait_for_elb_capacity}.
	// Experimental.
	WaitForElbCapacity *float64 `field:"optional" json:"waitForElbCapacity" yaml:"waitForElbCapacity"`
	// warm_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#warm_pool TfGroup#warm_pool}
	// Experimental.
	WarmPool *TfGroup_WarmPoolProperty `field:"optional" json:"warmPool" yaml:"warmPool"`
}

