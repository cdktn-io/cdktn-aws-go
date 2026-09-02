package awseventbridgepipes


// Experimental.
type TfPipe_EcsTaskParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#task_definition_arn TfPipe#task_definition_arn}.
	// Experimental.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// capacity_provider_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#capacity_provider_strategy TfPipe#capacity_provider_strategy}
	// Experimental.
	CapacityProviderStrategy interface{} `field:"optional" json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#enable_ecs_managed_tags TfPipe#enable_ecs_managed_tags}.
	// Experimental.
	EnableEcsManagedTags interface{} `field:"optional" json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#enable_execute_command TfPipe#enable_execute_command}.
	// Experimental.
	EnableExecuteCommand interface{} `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#group TfPipe#group}.
	// Experimental.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#launch_type TfPipe#launch_type}.
	// Experimental.
	LaunchType *string `field:"optional" json:"launchType" yaml:"launchType"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#network_configuration TfPipe#network_configuration}
	// Experimental.
	NetworkConfiguration *TfPipe_NetworkConfigurationProperty `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#overrides TfPipe#overrides}
	// Experimental.
	Overrides *TfPipe_OverridesProperty `field:"optional" json:"overrides" yaml:"overrides"`
	// placement_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#placement_constraint TfPipe#placement_constraint}
	// Experimental.
	PlacementConstraint interface{} `field:"optional" json:"placementConstraint" yaml:"placementConstraint"`
	// placement_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#placement_strategy TfPipe#placement_strategy}
	// Experimental.
	PlacementStrategy interface{} `field:"optional" json:"placementStrategy" yaml:"placementStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#platform_version TfPipe#platform_version}.
	// Experimental.
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#propagate_tags TfPipe#propagate_tags}.
	// Experimental.
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#reference_id TfPipe#reference_id}.
	// Experimental.
	ReferenceId *string `field:"optional" json:"referenceId" yaml:"referenceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#tags TfPipe#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#task_count TfPipe#task_count}.
	// Experimental.
	TaskCount *float64 `field:"optional" json:"taskCount" yaml:"taskCount"`
}

