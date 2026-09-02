package awsecs

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#name TfService#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// alarms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#alarms TfService#alarms}
	// Experimental.
	Alarms *TfService_AlarmsProperty `field:"optional" json:"alarms" yaml:"alarms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#availability_zone_rebalancing TfService#availability_zone_rebalancing}.
	// Experimental.
	AvailabilityZoneRebalancing *string `field:"optional" json:"availabilityZoneRebalancing" yaml:"availabilityZoneRebalancing"`
	// capacity_provider_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#capacity_provider_strategy TfService#capacity_provider_strategy}
	// Experimental.
	CapacityProviderStrategy interface{} `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#cluster TfService#cluster}.
	// Experimental.
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
	// deployment_circuit_breaker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#deployment_circuit_breaker TfService#deployment_circuit_breaker}
	// Experimental.
	DeploymentCircuitBreaker *TfService_DeploymentCircuitBreakerProperty `field:"optional" json:"deploymentCircuitBreaker" yaml:"deploymentCircuitBreaker"`
	// deployment_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#deployment_configuration TfService#deployment_configuration}
	// Experimental.
	DeploymentConfiguration *TfService_DeploymentConfigurationProperty `field:"optional" json:"deploymentConfiguration" yaml:"deploymentConfiguration"`
	// deployment_controller block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#deployment_controller TfService#deployment_controller}
	// Experimental.
	DeploymentController *TfService_DeploymentControllerProperty `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#deployment_maximum_percent TfService#deployment_maximum_percent}.
	// Experimental.
	DeploymentMaximumPercent *float64 `field:"optional" json:"deploymentMaximumPercent" yaml:"deploymentMaximumPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#deployment_minimum_healthy_percent TfService#deployment_minimum_healthy_percent}.
	// Experimental.
	DeploymentMinimumHealthyPercent *float64 `field:"optional" json:"deploymentMinimumHealthyPercent" yaml:"deploymentMinimumHealthyPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#desired_count TfService#desired_count}.
	// Experimental.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#enable_ecs_managed_tags TfService#enable_ecs_managed_tags}.
	// Experimental.
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#enable_execute_command TfService#enable_execute_command}.
	// Experimental.
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#force_delete TfService#force_delete}.
	// Experimental.
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#force_new_deployment TfService#force_new_deployment}.
	// Experimental.
	ForceNewDeployment interface{} `field:"optional" json:"forceNewDeployment" yaml:"forceNewDeployment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#health_check_grace_period_seconds TfService#health_check_grace_period_seconds}.
	// Experimental.
	HealthCheckGracePeriodSeconds *float64 `field:"optional" json:"healthCheckGracePeriodSeconds" yaml:"healthCheckGracePeriodSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#iam_role TfService#iam_role}.
	// Experimental.
	IamRole *string `field:"optional" json:"iamRole" yaml:"iamRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#id TfService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#launch_type TfService#launch_type}.
	// Experimental.
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#load_balancer TfService#load_balancer}
	// Experimental.
	LoadBalancer interface{} `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#network_configuration TfService#network_configuration}
	// Experimental.
	NetworkConfiguration *TfService_NetworkConfigurationProperty `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// ordered_placement_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#ordered_placement_strategy TfService#ordered_placement_strategy}
	// Experimental.
	OrderedPlacementStrategy interface{} `field:"optional" json:"orderedPlacementStrategy" yaml:"orderedPlacementStrategy"`
	// placement_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#placement_constraints TfService#placement_constraints}
	// Experimental.
	PlacementConstraints interface{} `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#platform_version TfService#platform_version}.
	// Experimental.
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#propagate_tags TfService#propagate_tags}.
	// Experimental.
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#region TfService#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#scheduling_strategy TfService#scheduling_strategy}.
	// Experimental.
	SchedulingStrategy *string `field:"optional" json:"schedulingStrategy" yaml:"schedulingStrategy"`
	// service_connect_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#service_connect_configuration TfService#service_connect_configuration}
	// Experimental.
	ServiceConnectConfiguration *TfService_ServiceConnectConfigurationProperty `field:"optional" json:"serviceConnectConfiguration" yaml:"serviceConnectConfiguration"`
	// service_registries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#service_registries TfService#service_registries}
	// Experimental.
	ServiceRegistries *TfService_ServiceRegistriesProperty `field:"optional" json:"serviceRegistries" yaml:"serviceRegistries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#sigint_rollback TfService#sigint_rollback}.
	// Experimental.
	SigintRollback interface{} `field:"optional" json:"sigintRollback" yaml:"sigintRollback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#tags TfService#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#tags_all TfService#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#task_definition TfService#task_definition}.
	// Experimental.
	TaskDefinition *string `field:"optional" json:"taskDefinition" yaml:"taskDefinition"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#timeouts TfService#timeouts}
	// Experimental.
	Timeouts *TfService_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#triggers TfService#triggers}.
	// Experimental.
	Triggers *map[string]*string `field:"optional" json:"triggers" yaml:"triggers"`
	// volume_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#volume_configuration TfService#volume_configuration}
	// Experimental.
	VolumeConfiguration *TfService_VolumeConfigurationProperty `field:"optional" json:"volumeConfiguration" yaml:"volumeConfiguration"`
	// vpc_lattice_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#vpc_lattice_configurations TfService#vpc_lattice_configurations}
	// Experimental.
	VpcLatticeConfigurations interface{} `field:"optional" json:"vpcLatticeConfigurations" yaml:"vpcLatticeConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#wait_for_steady_state TfService#wait_for_steady_state}.
	// Experimental.
	WaitForSteadyState interface{} `field:"optional" json:"waitForSteadyState" yaml:"waitForSteadyState"`
}

