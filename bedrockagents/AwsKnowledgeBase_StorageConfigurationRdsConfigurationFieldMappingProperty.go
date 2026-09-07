package bedrockagents


// Experimental.
type AwsKnowledgeBase_StorageConfigurationRdsConfigurationFieldMappingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#metadata_field AwsKnowledgeBase#metadata_field}.
	// Experimental.
	MetadataField *string `field:"required" json:"metadataField" yaml:"metadataField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#primary_key_field AwsKnowledgeBase#primary_key_field}.
	// Experimental.
	PrimaryKeyField *string `field:"required" json:"primaryKeyField" yaml:"primaryKeyField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#text_field AwsKnowledgeBase#text_field}.
	// Experimental.
	TextField *string `field:"required" json:"textField" yaml:"textField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#vector_field AwsKnowledgeBase#vector_field}.
	// Experimental.
	VectorField *string `field:"required" json:"vectorField" yaml:"vectorField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#custom_metadata_field AwsKnowledgeBase#custom_metadata_field}.
	// Experimental.
	CustomMetadataField *string `field:"optional" json:"customMetadataField" yaml:"customMetadataField"`
}

