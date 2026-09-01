package awsecs


// Experimental.
type AwsEcsCapacityProvider_ManagedInstancesProviderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#infrastructure_role_arn AwsEcsCapacityProvider#infrastructure_role_arn}.
	// Experimental.
	InfrastructureRoleArn *string `field:"required" json:"infrastructureRoleArn" yaml:"infrastructureRoleArn"`
	// instance_launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#instance_launch_template AwsEcsCapacityProvider#instance_launch_template}
	// Experimental.
	InstanceLaunchTemplate *AwsEcsCapacityProvider_InstanceLaunchTemplateProperty `field:"required" json:"instanceLaunchTemplate" yaml:"instanceLaunchTemplate"`
	// infrastructure_optimization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#infrastructure_optimization AwsEcsCapacityProvider#infrastructure_optimization}
	// Experimental.
	InfrastructureOptimization *AwsEcsCapacityProvider_InfrastructureOptimizationProperty `field:"optional" json:"infrastructureOptimization" yaml:"infrastructureOptimization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#propagate_tags AwsEcsCapacityProvider#propagate_tags}.
	// Experimental.
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
}

