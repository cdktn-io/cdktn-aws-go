package awseventbridge


// Experimental.
type TfTarget_EcsTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#task_definition_arn TfTarget#task_definition_arn}.
	// Experimental.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// capacity_provider_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#capacity_provider_strategy TfTarget#capacity_provider_strategy}
	// Experimental.
	CapacityProviderStrategy interface{} `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#enable_ecs_managed_tags TfTarget#enable_ecs_managed_tags}.
	// Experimental.
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#enable_execute_command TfTarget#enable_execute_command}.
	// Experimental.
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#group TfTarget#group}.
	// Experimental.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#launch_type TfTarget#launch_type}.
	// Experimental.
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#network_configuration TfTarget#network_configuration}
	// Experimental.
	NetworkConfiguration *TfTarget_NetworkConfigurationProperty `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// ordered_placement_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#ordered_placement_strategy TfTarget#ordered_placement_strategy}
	// Experimental.
	OrderedPlacementStrategy interface{} `field:"optional" json:"orderedPlacementStrategy" yaml:"orderedPlacementStrategy"`
	// placement_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#placement_constraint TfTarget#placement_constraint}
	// Experimental.
	PlacementConstraint interface{} `field:"optional" json:"placementConstraint" yaml:"placementConstraint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#platform_version TfTarget#platform_version}.
	// Experimental.
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#propagate_tags TfTarget#propagate_tags}.
	// Experimental.
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#tags TfTarget#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#task_count TfTarget#task_count}.
	// Experimental.
	TaskCount *float64 `field:"optional" json:"taskCount" yaml:"taskCount"`
}

