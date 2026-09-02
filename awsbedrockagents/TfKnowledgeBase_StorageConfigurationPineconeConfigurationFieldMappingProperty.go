package awsbedrockagents


// Experimental.
type TfKnowledgeBase_StorageConfigurationPineconeConfigurationFieldMappingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#metadata_field TfKnowledgeBase#metadata_field}.
	// Experimental.
	MetadataField *string `field:"required" json:"metadataField" yaml:"metadataField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#text_field TfKnowledgeBase#text_field}.
	// Experimental.
	TextField *string `field:"required" json:"textField" yaml:"textField"`
}

