package awscloudwatchobservabilityadmin


// Experimental.
type AwsObservabilityadminTelemetryRule_WafLoggingParametersProperty struct {
	// logging_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#logging_filter AwsObservabilityadminTelemetryRule#logging_filter}
	// Experimental.
	LoggingFilter interface{} `field:"optional" json:"loggingFilter" yaml:"loggingFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#log_type AwsObservabilityadminTelemetryRule#log_type}.
	// Experimental.
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
	// redacted_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#redacted_fields AwsObservabilityadminTelemetryRule#redacted_fields}
	// Experimental.
	RedactedFields interface{} `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

