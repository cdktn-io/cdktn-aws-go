package lambda


// Experimental.
type AwsCapacityProvider_ScalingPoliciesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_capacity_provider#predefined_metric_type AwsCapacityProvider#predefined_metric_type}.
	// Experimental.
	PredefinedMetricType *string `field:"optional" json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_capacity_provider#target_value AwsCapacityProvider#target_value}.
	// Experimental.
	TargetValue *float64 `field:"optional" json:"targetValue" yaml:"targetValue"`
}

