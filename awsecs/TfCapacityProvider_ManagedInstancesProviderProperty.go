package awsecs


// Experimental.
type TfCapacityProvider_ManagedInstancesProviderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#infrastructure_role_arn TfCapacityProvider#infrastructure_role_arn}.
	// Experimental.
	InfrastructureRoleArn *string `field:"required" json:"infrastructureRoleArn" yaml:"infrastructureRoleArn"`
	// instance_launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#instance_launch_template TfCapacityProvider#instance_launch_template}
	// Experimental.
	InstanceLaunchTemplate *TfCapacityProvider_InstanceLaunchTemplateProperty `field:"required" json:"instanceLaunchTemplate" yaml:"instanceLaunchTemplate"`
	// infrastructure_optimization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#infrastructure_optimization TfCapacityProvider#infrastructure_optimization}
	// Experimental.
	InfrastructureOptimization *TfCapacityProvider_InfrastructureOptimizationProperty `field:"optional" json:"infrastructureOptimization" yaml:"infrastructureOptimization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#propagate_tags TfCapacityProvider#propagate_tags}.
	// Experimental.
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
}

