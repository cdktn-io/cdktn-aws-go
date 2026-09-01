package awsbedrockagents


// Experimental.
type AwsBedrockagentKnowledgeBase_StorageConfigurationRedisEnterpriseCloudConfigurationFieldMappingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#metadata_field AwsBedrockagentKnowledgeBase#metadata_field}.
	// Experimental.
	MetadataField *string `field:"optional" json:"metadataField" yaml:"metadataField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#text_field AwsBedrockagentKnowledgeBase#text_field}.
	// Experimental.
	TextField *string `field:"optional" json:"textField" yaml:"textField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#vector_field AwsBedrockagentKnowledgeBase#vector_field}.
	// Experimental.
	VectorField *string `field:"optional" json:"vectorField" yaml:"vectorField"`
}

