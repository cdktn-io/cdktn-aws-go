package awscloudwatchobservabilityadmin


// Experimental.
type TfTelemetryRuleForOrganization_WafLoggingParametersProperty struct {
	// logging_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#logging_filter TfTelemetryRuleForOrganization#logging_filter}
	// Experimental.
	LoggingFilter interface{} `field:"optional" json:"loggingFilter" yaml:"loggingFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#log_type TfTelemetryRuleForOrganization#log_type}.
	// Experimental.
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
	// redacted_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#redacted_fields TfTelemetryRuleForOrganization#redacted_fields}
	// Experimental.
	RedactedFields interface{} `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

