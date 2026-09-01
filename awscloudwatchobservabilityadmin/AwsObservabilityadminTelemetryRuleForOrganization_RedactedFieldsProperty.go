package awscloudwatchobservabilityadmin


// Experimental.
type AwsObservabilityadminTelemetryRuleForOrganization_RedactedFieldsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#method AwsObservabilityadminTelemetryRuleForOrganization#method}.
	// Experimental.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#query_string AwsObservabilityadminTelemetryRuleForOrganization#query_string}.
	// Experimental.
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// single_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#single_header AwsObservabilityadminTelemetryRuleForOrganization#single_header}
	// Experimental.
	SingleHeader interface{} `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#uri_path AwsObservabilityadminTelemetryRuleForOrganization#uri_path}.
	// Experimental.
	UriPath *string `field:"optional" json:"uriPath" yaml:"uriPath"`
}

