package awscloudwatchobservabilityadmin


// Experimental.
type TfCentralizationRuleForOrganization_DestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#account TfCentralizationRuleForOrganization#account}.
	// Experimental.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#region TfCentralizationRuleForOrganization#region}.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
	// destination_logs_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#destination_logs_configuration TfCentralizationRuleForOrganization#destination_logs_configuration}
	// Experimental.
	DestinationLogsConfiguration interface{} `field:"optional" json:"destinationLogsConfiguration" yaml:"destinationLogsConfiguration"`
	// destination_metrics_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#destination_metrics_configuration TfCentralizationRuleForOrganization#destination_metrics_configuration}
	// Experimental.
	DestinationMetricsConfiguration interface{} `field:"optional" json:"destinationMetricsConfiguration" yaml:"destinationMetricsConfiguration"`
}

