package ecs


// Experimental.
type AwsService_DeploymentConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#bake_time_in_minutes AwsService#bake_time_in_minutes}.
	// Experimental.
	BakeTimeInMinutes *string `field:"optional" json:"bakeTimeInMinutes" yaml:"bakeTimeInMinutes"`
	// canary_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#canary_configuration AwsService#canary_configuration}
	// Experimental.
	CanaryConfiguration *AwsService_CanaryConfigurationProperty `field:"optional" json:"canaryConfiguration" yaml:"canaryConfiguration"`
	// lifecycle_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#lifecycle_hook AwsService#lifecycle_hook}
	// Experimental.
	LifecycleHook interface{} `field:"optional" json:"lifecycleHook" yaml:"lifecycleHook"`
	// linear_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#linear_configuration AwsService#linear_configuration}
	// Experimental.
	LinearConfiguration *AwsService_LinearConfigurationProperty `field:"optional" json:"linearConfiguration" yaml:"linearConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#strategy AwsService#strategy}.
	// Experimental.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

