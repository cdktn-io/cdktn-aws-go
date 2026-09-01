package awsbedrockagents


// Experimental.
type AwsBedrockagentKnowledgeBase_KnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationEmbeddingModelConfigurationBedrockEmbeddingModelConfigurationProperty struct {
	// audio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#audio AwsBedrockagentKnowledgeBase#audio}
	// Experimental.
	Audio interface{} `field:"optional" json:"audio" yaml:"audio"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#dimensions AwsBedrockagentKnowledgeBase#dimensions}.
	// Experimental.
	Dimensions *float64 `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#embedding_data_type AwsBedrockagentKnowledgeBase#embedding_data_type}.
	// Experimental.
	EmbeddingDataType *string `field:"optional" json:"embeddingDataType" yaml:"embeddingDataType"`
	// video block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#video AwsBedrockagentKnowledgeBase#video}
	// Experimental.
	Video interface{} `field:"optional" json:"video" yaml:"video"`
}

