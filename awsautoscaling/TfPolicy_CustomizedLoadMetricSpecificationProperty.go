package awsautoscaling


// Experimental.
type TfPolicy_CustomizedLoadMetricSpecificationProperty struct {
	// metric_data_queries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#metric_data_queries TfPolicy#metric_data_queries}
	// Experimental.
	MetricDataQueries interface{} `field:"required" json:"metricDataQueries" yaml:"metricDataQueries"`
}

