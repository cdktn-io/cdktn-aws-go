package cloudwatchobservabilityadmin


// Experimental.
type AwsTelemetryRuleForOrganization_AdvancedEventSelectorsProperty struct {
	// field_selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#field_selectors AwsTelemetryRuleForOrganization#field_selectors}
	// Experimental.
	FieldSelectors interface{} `field:"optional" json:"fieldSelectors" yaml:"fieldSelectors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#name AwsTelemetryRuleForOrganization#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

