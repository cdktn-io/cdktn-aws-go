package awsapplicationautoscaling


// Experimental.
type TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#metric_name TfPolicy#metric_name}.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#namespace TfPolicy#namespace}.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#dimensions TfPolicy#dimensions}
	// Experimental.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
}

