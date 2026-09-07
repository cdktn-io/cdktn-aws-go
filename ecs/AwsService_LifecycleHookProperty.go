package ecs


// Experimental.
type AwsService_LifecycleHookProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#hook_target_arn AwsService#hook_target_arn}.
	// Experimental.
	HookTargetArn *string `field:"required" json:"hookTargetArn" yaml:"hookTargetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#lifecycle_stages AwsService#lifecycle_stages}.
	// Experimental.
	LifecycleStages *[]*string `field:"required" json:"lifecycleStages" yaml:"lifecycleStages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#role_arn AwsService#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#hook_details AwsService#hook_details}.
	// Experimental.
	HookDetails *string `field:"optional" json:"hookDetails" yaml:"hookDetails"`
}

