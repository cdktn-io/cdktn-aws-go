package securityhub


// Experimental.
type AwsConnectorV2_ConnectorProviderProperty struct {
	// jira_cloud block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_connector_v2#jira_cloud AwsConnectorV2#jira_cloud}
	// Experimental.
	JiraCloud interface{} `field:"optional" json:"jiraCloud" yaml:"jiraCloud"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_connector_v2#service_now AwsConnectorV2#service_now}
	// Experimental.
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
}

