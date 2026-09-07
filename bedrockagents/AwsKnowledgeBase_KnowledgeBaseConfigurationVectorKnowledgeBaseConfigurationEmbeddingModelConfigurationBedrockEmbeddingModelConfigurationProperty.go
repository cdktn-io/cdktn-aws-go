package bedrockagents


// Experimental.
type AwsKnowledgeBase_KnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationEmbeddingModelConfigurationBedrockEmbeddingModelConfigurationProperty struct {
	// audio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#audio AwsKnowledgeBase#audio}
	// Experimental.
	Audio interface{} `field:"optional" json:"audio" yaml:"audio"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#dimensions AwsKnowledgeBase#dimensions}.
	// Experimental.
	Dimensions *float64 `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#embedding_data_type AwsKnowledgeBase#embedding_data_type}.
	// Experimental.
	EmbeddingDataType *string `field:"optional" json:"embeddingDataType" yaml:"embeddingDataType"`
	// video block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#video AwsKnowledgeBase#video}
	// Experimental.
	Video interface{} `field:"optional" json:"video" yaml:"video"`
}

