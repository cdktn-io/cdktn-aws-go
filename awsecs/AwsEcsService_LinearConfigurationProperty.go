package awsecs


// Experimental.
type AwsEcsService_LinearConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#step_bake_time_in_minutes AwsEcsService#step_bake_time_in_minutes}.
	// Experimental.
	StepBakeTimeInMinutes *string `field:"optional" json:"stepBakeTimeInMinutes" yaml:"stepBakeTimeInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#step_percent AwsEcsService#step_percent}.
	// Experimental.
	StepPercent *float64 `field:"optional" json:"stepPercent" yaml:"stepPercent"`
}

