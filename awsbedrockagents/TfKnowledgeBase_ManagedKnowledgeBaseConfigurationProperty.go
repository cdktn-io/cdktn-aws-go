package awsbedrockagents


// Experimental.
type TfKnowledgeBase_ManagedKnowledgeBaseConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#embedding_model_arn TfKnowledgeBase#embedding_model_arn}.
	// Experimental.
	EmbeddingModelArn *string `field:"optional" json:"embeddingModelArn" yaml:"embeddingModelArn"`
	// embedding_model_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#embedding_model_configuration TfKnowledgeBase#embedding_model_configuration}
	// Experimental.
	EmbeddingModelConfiguration interface{} `field:"optional" json:"embeddingModelConfiguration" yaml:"embeddingModelConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#embedding_model_type TfKnowledgeBase#embedding_model_type}.
	// Experimental.
	EmbeddingModelType *string `field:"optional" json:"embeddingModelType" yaml:"embeddingModelType"`
	// server_side_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#server_side_encryption_configuration TfKnowledgeBase#server_side_encryption_configuration}
	// Experimental.
	ServerSideEncryptionConfiguration interface{} `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
}

