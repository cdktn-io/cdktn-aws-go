package awscloudwatchobservabilityadmin


// Experimental.
type TfTelemetryRule_FieldSelectorsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#field TfTelemetryRule#field}.
	// Experimental.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#ends_with TfTelemetryRule#ends_with}.
	// Experimental.
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#equals TfTelemetryRule#equals}.
	// Experimental.
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#not_ends_with TfTelemetryRule#not_ends_with}.
	// Experimental.
	NotEndsWith *[]*string `field:"optional" json:"notEndsWith" yaml:"notEndsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#not_equals TfTelemetryRule#not_equals}.
	// Experimental.
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#not_starts_with TfTelemetryRule#not_starts_with}.
	// Experimental.
	NotStartsWith *[]*string `field:"optional" json:"notStartsWith" yaml:"notStartsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#starts_with TfTelemetryRule#starts_with}.
	// Experimental.
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

