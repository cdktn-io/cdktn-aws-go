package awscodebuild


// Experimental.
type AwsCodebuildFleet_ScalingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#max_capacity AwsCodebuildFleet#max_capacity}.
	// Experimental.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#scaling_type AwsCodebuildFleet#scaling_type}.
	// Experimental.
	ScalingType *string `field:"optional" json:"scalingType" yaml:"scalingType"`
	// target_tracking_scaling_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#target_tracking_scaling_configs AwsCodebuildFleet#target_tracking_scaling_configs}
	// Experimental.
	TargetTrackingScalingConfigs interface{} `field:"optional" json:"targetTrackingScalingConfigs" yaml:"targetTrackingScalingConfigs"`
}

