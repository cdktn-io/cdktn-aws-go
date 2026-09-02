package awsapplicationautoscaling


// Experimental.
type TfPolicy_PredictiveScalingPolicyConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricProperty struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#dimension TfPolicy#dimension}
	// Experimental.
	Dimension interface{} `field:"optional" json:"dimension" yaml:"dimension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#metric_name TfPolicy#metric_name}.
	// Experimental.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#namespace TfPolicy#namespace}.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

