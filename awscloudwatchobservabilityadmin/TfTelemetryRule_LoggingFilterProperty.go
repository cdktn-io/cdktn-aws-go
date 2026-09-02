package awscloudwatchobservabilityadmin


// Experimental.
type TfTelemetryRule_LoggingFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#default_behavior TfTelemetryRule#default_behavior}.
	// Experimental.
	DefaultBehavior *string `field:"optional" json:"defaultBehavior" yaml:"defaultBehavior"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#filters TfTelemetryRule#filters}
	// Experimental.
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

