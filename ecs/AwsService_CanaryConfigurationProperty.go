package ecs


// Experimental.
type AwsService_CanaryConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#canary_bake_time_in_minutes AwsService#canary_bake_time_in_minutes}.
	// Experimental.
	CanaryBakeTimeInMinutes *string `field:"optional" json:"canaryBakeTimeInMinutes" yaml:"canaryBakeTimeInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#canary_percent AwsService#canary_percent}.
	// Experimental.
	CanaryPercent *float64 `field:"optional" json:"canaryPercent" yaml:"canaryPercent"`
}

