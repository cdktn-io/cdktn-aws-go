package awsamp


// Experimental.
type TfQueryLoggingConfiguration_CloudwatchLogsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_query_logging_configuration#log_group_arn TfQueryLoggingConfiguration#log_group_arn}.
	// Experimental.
	LogGroupArn *string `field:"required" json:"logGroupArn" yaml:"logGroupArn"`
}

