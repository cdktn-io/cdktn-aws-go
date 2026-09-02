package awsamp


// Experimental.
type TfQueryLoggingConfiguration_DestinationProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_query_logging_configuration#cloudwatch_logs TfQueryLoggingConfiguration#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs interface{} `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_query_logging_configuration#filters TfQueryLoggingConfiguration#filters}
	// Experimental.
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

