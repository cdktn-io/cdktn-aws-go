package awsautoscaling


// Experimental.
type AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#name AwsAutoscalingPolicy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#value AwsAutoscalingPolicy#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

