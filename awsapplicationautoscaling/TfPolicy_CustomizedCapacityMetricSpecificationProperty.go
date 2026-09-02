package awsapplicationautoscaling


// Experimental.
type TfPolicy_CustomizedCapacityMetricSpecificationProperty struct {
	// metric_data_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appautoscaling_policy#metric_data_query TfPolicy#metric_data_query}
	// Experimental.
	MetricDataQuery interface{} `field:"required" json:"metricDataQuery" yaml:"metricDataQuery"`
}

