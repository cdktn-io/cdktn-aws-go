package awsautoscaling


// Experimental.
type AwsAutoscalingPolicy_CustomizedCapacityMetricSpecificationProperty struct {
	// metric_data_queries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#metric_data_queries AwsAutoscalingPolicy#metric_data_queries}
	// Experimental.
	MetricDataQueries interface{} `field:"required" json:"metricDataQueries" yaml:"metricDataQueries"`
}

