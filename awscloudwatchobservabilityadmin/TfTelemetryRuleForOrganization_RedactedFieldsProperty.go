package awscloudwatchobservabilityadmin


// Experimental.
type TfTelemetryRuleForOrganization_RedactedFieldsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#method TfTelemetryRuleForOrganization#method}.
	// Experimental.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#query_string TfTelemetryRuleForOrganization#query_string}.
	// Experimental.
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// single_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#single_header TfTelemetryRuleForOrganization#single_header}
	// Experimental.
	SingleHeader interface{} `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#uri_path TfTelemetryRuleForOrganization#uri_path}.
	// Experimental.
	UriPath *string `field:"optional" json:"uriPath" yaml:"uriPath"`
}

