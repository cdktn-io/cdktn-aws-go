package awsecs


// Experimental.
type TfService_LifecycleHookProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#hook_target_arn TfService#hook_target_arn}.
	// Experimental.
	HookTargetArn *string `field:"required" json:"hookTargetArn" yaml:"hookTargetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#lifecycle_stages TfService#lifecycle_stages}.
	// Experimental.
	LifecycleStages *[]*string `field:"required" json:"lifecycleStages" yaml:"lifecycleStages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#role_arn TfService#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#hook_details TfService#hook_details}.
	// Experimental.
	HookDetails *string `field:"optional" json:"hookDetails" yaml:"hookDetails"`
}

