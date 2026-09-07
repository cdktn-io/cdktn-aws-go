package bedrockagents


// Experimental.
type AwsDataSource_ManagedKnowledgeBaseConnectorConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#connector_parameters AwsDataSource#connector_parameters}.
	// Experimental.
	ConnectorParameters *string `field:"optional" json:"connectorParameters" yaml:"connectorParameters"`
	// deletion_protection_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#deletion_protection_configuration AwsDataSource#deletion_protection_configuration}
	// Experimental.
	DeletionProtectionConfiguration interface{} `field:"optional" json:"deletionProtectionConfiguration" yaml:"deletionProtectionConfiguration"`
	// media_extraction_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#media_extraction_configuration AwsDataSource#media_extraction_configuration}
	// Experimental.
	MediaExtractionConfiguration interface{} `field:"optional" json:"mediaExtractionConfiguration" yaml:"mediaExtractionConfiguration"`
}

