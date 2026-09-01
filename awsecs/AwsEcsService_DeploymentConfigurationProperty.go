package awsecs


// Experimental.
type AwsEcsService_DeploymentConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#bake_time_in_minutes AwsEcsService#bake_time_in_minutes}.
	// Experimental.
	BakeTimeInMinutes *string `field:"optional" json:"bakeTimeInMinutes" yaml:"bakeTimeInMinutes"`
	// canary_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#canary_configuration AwsEcsService#canary_configuration}
	// Experimental.
	CanaryConfiguration *AwsEcsService_CanaryConfigurationProperty `field:"optional" json:"canaryConfiguration" yaml:"canaryConfiguration"`
	// lifecycle_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#lifecycle_hook AwsEcsService#lifecycle_hook}
	// Experimental.
	LifecycleHook interface{} `field:"optional" json:"lifecycleHook" yaml:"lifecycleHook"`
	// linear_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#linear_configuration AwsEcsService#linear_configuration}
	// Experimental.
	LinearConfiguration *AwsEcsService_LinearConfigurationProperty `field:"optional" json:"linearConfiguration" yaml:"linearConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#strategy AwsEcsService#strategy}.
	// Experimental.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

