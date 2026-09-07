package autoscaling


// Experimental.
type AwsPolicy_CustomizedCapacityMetricSpecificationProperty struct {
	// metric_data_queries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_policy#metric_data_queries AwsPolicy#metric_data_queries}
	// Experimental.
	MetricDataQueries interface{} `field:"required" json:"metricDataQueries" yaml:"metricDataQueries"`
}

