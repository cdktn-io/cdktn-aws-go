package awsecs


// Experimental.
type AwsEcsService_CanaryConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#canary_bake_time_in_minutes AwsEcsService#canary_bake_time_in_minutes}.
	// Experimental.
	CanaryBakeTimeInMinutes *string `field:"optional" json:"canaryBakeTimeInMinutes" yaml:"canaryBakeTimeInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#canary_percent AwsEcsService#canary_percent}.
	// Experimental.
	CanaryPercent *float64 `field:"optional" json:"canaryPercent" yaml:"canaryPercent"`
}

