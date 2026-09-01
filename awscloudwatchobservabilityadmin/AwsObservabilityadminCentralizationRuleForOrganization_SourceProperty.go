package awscloudwatchobservabilityadmin


// Experimental.
type AwsObservabilityadminCentralizationRuleForOrganization_SourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#regions AwsObservabilityadminCentralizationRuleForOrganization#regions}.
	// Experimental.
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#scope AwsObservabilityadminCentralizationRuleForOrganization#scope}.
	// Experimental.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// source_logs_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#source_logs_configuration AwsObservabilityadminCentralizationRuleForOrganization#source_logs_configuration}
	// Experimental.
	SourceLogsConfiguration interface{} `field:"optional" json:"sourceLogsConfiguration" yaml:"sourceLogsConfiguration"`
	// source_metrics_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#source_metrics_configuration AwsObservabilityadminCentralizationRuleForOrganization#source_metrics_configuration}
	// Experimental.
	SourceMetricsConfiguration interface{} `field:"optional" json:"sourceMetricsConfiguration" yaml:"sourceMetricsConfiguration"`
}

