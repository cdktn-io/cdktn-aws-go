package awssecurityhub


// Experimental.
type TfAutomationRuleV2_ExternalIntegrationConfigurationProperty struct {
	// The ARN of the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule_v2#connector_arn TfAutomationRuleV2#connector_arn}
	// Experimental.
	ConnectorArn *string `field:"required" json:"connectorArn" yaml:"connectorArn"`
}

