package awscloudwatchobservabilityadmin


// Experimental.
type TfTelemetryRuleForOrganization_FiltersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#behavior TfTelemetryRuleForOrganization#behavior}.
	// Experimental.
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#conditions TfTelemetryRuleForOrganization#conditions}
	// Experimental.
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#requirement TfTelemetryRuleForOrganization#requirement}.
	// Experimental.
	Requirement *string `field:"optional" json:"requirement" yaml:"requirement"`
}

