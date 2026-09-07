package cloudwatchevidently


// Experimental.
type AwsLaunch_MetricMonitorsProperty struct {
	// metric_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#metric_definition AwsLaunch#metric_definition}
	// Experimental.
	MetricDefinition *AwsLaunch_MetricDefinitionProperty `field:"required" json:"metricDefinition" yaml:"metricDefinition"`
}

