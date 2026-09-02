package awsbedrockagents


// Experimental.
type TfKnowledgeBase_VectorKnowledgeBaseConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#embedding_model_arn TfKnowledgeBase#embedding_model_arn}.
	// Experimental.
	EmbeddingModelArn *string `field:"required" json:"embeddingModelArn" yaml:"embeddingModelArn"`
	// embedding_model_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#embedding_model_configuration TfKnowledgeBase#embedding_model_configuration}
	// Experimental.
	EmbeddingModelConfiguration interface{} `field:"optional" json:"embeddingModelConfiguration" yaml:"embeddingModelConfiguration"`
	// supplemental_data_storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#supplemental_data_storage_configuration TfKnowledgeBase#supplemental_data_storage_configuration}
	// Experimental.
	SupplementalDataStorageConfiguration interface{} `field:"optional" json:"supplementalDataStorageConfiguration" yaml:"supplementalDataStorageConfiguration"`
}

