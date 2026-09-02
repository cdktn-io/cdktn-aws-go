package awsautoscaling


// Experimental.
type TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#metric_name TfPolicy#metric_name}.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#namespace TfPolicy#namespace}.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#dimensions TfPolicy#dimensions}
	// Experimental.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
}

