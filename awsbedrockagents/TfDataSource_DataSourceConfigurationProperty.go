package awsbedrockagents


// Experimental.
type TfDataSource_DataSourceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#type TfDataSource#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// confluence_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#confluence_configuration TfDataSource#confluence_configuration}
	// Experimental.
	ConfluenceConfiguration interface{} `field:"optional" json:"confluenceConfiguration" yaml:"confluenceConfiguration"`
	// managed_knowledge_base_connector_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#managed_knowledge_base_connector_configuration TfDataSource#managed_knowledge_base_connector_configuration}
	// Experimental.
	ManagedKnowledgeBaseConnectorConfiguration interface{} `field:"optional" json:"managedKnowledgeBaseConnectorConfiguration" yaml:"managedKnowledgeBaseConnectorConfiguration"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#s3_configuration TfDataSource#s3_configuration}
	// Experimental.
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
	// salesforce_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#salesforce_configuration TfDataSource#salesforce_configuration}
	// Experimental.
	SalesforceConfiguration interface{} `field:"optional" json:"salesforceConfiguration" yaml:"salesforceConfiguration"`
	// share_point_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#share_point_configuration TfDataSource#share_point_configuration}
	// Experimental.
	SharePointConfiguration interface{} `field:"optional" json:"sharePointConfiguration" yaml:"sharePointConfiguration"`
	// web_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#web_configuration TfDataSource#web_configuration}
	// Experimental.
	WebConfiguration interface{} `field:"optional" json:"webConfiguration" yaml:"webConfiguration"`
}

