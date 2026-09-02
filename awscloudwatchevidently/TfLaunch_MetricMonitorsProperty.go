package awscloudwatchevidently


// Experimental.
type TfLaunch_MetricMonitorsProperty struct {
	// metric_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#metric_definition TfLaunch#metric_definition}
	// Experimental.
	MetricDefinition *TfLaunch_MetricDefinitionProperty `field:"required" json:"metricDefinition" yaml:"metricDefinition"`
}

