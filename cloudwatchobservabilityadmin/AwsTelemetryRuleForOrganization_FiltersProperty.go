package cloudwatchobservabilityadmin


// Experimental.
type AwsTelemetryRuleForOrganization_FiltersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#behavior AwsTelemetryRuleForOrganization#behavior}.
	// Experimental.
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#conditions AwsTelemetryRuleForOrganization#conditions}
	// Experimental.
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#requirement AwsTelemetryRuleForOrganization#requirement}.
	// Experimental.
	Requirement *string `field:"optional" json:"requirement" yaml:"requirement"`
}

