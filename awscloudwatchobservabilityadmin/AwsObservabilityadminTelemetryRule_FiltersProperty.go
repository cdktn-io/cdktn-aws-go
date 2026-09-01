package awscloudwatchobservabilityadmin


// Experimental.
type AwsObservabilityadminTelemetryRule_FiltersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#behavior AwsObservabilityadminTelemetryRule#behavior}.
	// Experimental.
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#conditions AwsObservabilityadminTelemetryRule#conditions}
	// Experimental.
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#requirement AwsObservabilityadminTelemetryRule#requirement}.
	// Experimental.
	Requirement *string `field:"optional" json:"requirement" yaml:"requirement"`
}

