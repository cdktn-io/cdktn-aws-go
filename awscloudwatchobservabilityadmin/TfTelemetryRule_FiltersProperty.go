package awscloudwatchobservabilityadmin


// Experimental.
type TfTelemetryRule_FiltersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#behavior TfTelemetryRule#behavior}.
	// Experimental.
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#conditions TfTelemetryRule#conditions}
	// Experimental.
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#requirement TfTelemetryRule#requirement}.
	// Experimental.
	Requirement *string `field:"optional" json:"requirement" yaml:"requirement"`
}

