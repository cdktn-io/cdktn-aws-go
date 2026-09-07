package bedrockagents


// Experimental.
type AwsKnowledgeBase_StorageConfigurationPineconeConfigurationFieldMappingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#metadata_field AwsKnowledgeBase#metadata_field}.
	// Experimental.
	MetadataField *string `field:"required" json:"metadataField" yaml:"metadataField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#text_field AwsKnowledgeBase#text_field}.
	// Experimental.
	TextField *string `field:"required" json:"textField" yaml:"textField"`
}

