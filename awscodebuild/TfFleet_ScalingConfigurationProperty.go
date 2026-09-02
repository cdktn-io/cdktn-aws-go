package awscodebuild


// Experimental.
type TfFleet_ScalingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#max_capacity TfFleet#max_capacity}.
	// Experimental.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#scaling_type TfFleet#scaling_type}.
	// Experimental.
	ScalingType *string `field:"optional" json:"scalingType" yaml:"scalingType"`
	// target_tracking_scaling_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#target_tracking_scaling_configs TfFleet#target_tracking_scaling_configs}
	// Experimental.
	TargetTrackingScalingConfigs interface{} `field:"optional" json:"targetTrackingScalingConfigs" yaml:"targetTrackingScalingConfigs"`
}

