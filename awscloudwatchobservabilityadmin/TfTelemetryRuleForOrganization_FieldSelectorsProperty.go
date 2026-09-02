package awscloudwatchobservabilityadmin


// Experimental.
type TfTelemetryRuleForOrganization_FieldSelectorsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#field TfTelemetryRuleForOrganization#field}.
	// Experimental.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#ends_with TfTelemetryRuleForOrganization#ends_with}.
	// Experimental.
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#equals TfTelemetryRuleForOrganization#equals}.
	// Experimental.
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#not_ends_with TfTelemetryRuleForOrganization#not_ends_with}.
	// Experimental.
	NotEndsWith *[]*string `field:"optional" json:"notEndsWith" yaml:"notEndsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#not_equals TfTelemetryRuleForOrganization#not_equals}.
	// Experimental.
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#not_starts_with TfTelemetryRuleForOrganization#not_starts_with}.
	// Experimental.
	NotStartsWith *[]*string `field:"optional" json:"notStartsWith" yaml:"notStartsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#starts_with TfTelemetryRuleForOrganization#starts_with}.
	// Experimental.
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

