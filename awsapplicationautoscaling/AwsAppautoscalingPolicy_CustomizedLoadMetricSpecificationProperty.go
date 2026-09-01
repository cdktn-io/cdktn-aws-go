package awsapplicationautoscaling


// Experimental.
type AwsAppautoscalingPolicy_CustomizedLoadMetricSpecificationProperty struct {
	// metric_data_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#metric_data_query AwsAppautoscalingPolicy#metric_data_query}
	// Experimental.
	MetricDataQuery interface{} `field:"required" json:"metricDataQuery" yaml:"metricDataQuery"`
}

