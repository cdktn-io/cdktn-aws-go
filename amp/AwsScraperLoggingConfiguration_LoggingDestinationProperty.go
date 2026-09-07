package amp


// Experimental.
type AwsScraperLoggingConfiguration_LoggingDestinationProperty struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper_logging_configuration#cloudwatch_logs AwsScraperLoggingConfiguration#cloudwatch_logs}
	// Experimental.
	CloudwatchLogs interface{} `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
}

