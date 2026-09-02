package awscloudwatchobservabilityadmin


// Experimental.
type TfTelemetryRuleForOrganization_VpcFlowLogParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#log_format TfTelemetryRuleForOrganization#log_format}.
	// Experimental.
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#max_aggregation_interval TfTelemetryRuleForOrganization#max_aggregation_interval}.
	// Experimental.
	MaxAggregationInterval *float64 `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#traffic_type TfTelemetryRuleForOrganization#traffic_type}.
	// Experimental.
	TrafficType *string `field:"optional" json:"trafficType" yaml:"trafficType"`
}

