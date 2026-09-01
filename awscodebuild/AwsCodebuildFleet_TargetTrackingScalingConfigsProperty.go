package awscodebuild


// Experimental.
type AwsCodebuildFleet_TargetTrackingScalingConfigsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#metric_type AwsCodebuildFleet#metric_type}.
	// Experimental.
	MetricType *string `field:"optional" json:"metricType" yaml:"metricType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_fleet#target_value AwsCodebuildFleet#target_value}.
	// Experimental.
	TargetValue *float64 `field:"optional" json:"targetValue" yaml:"targetValue"`
}

