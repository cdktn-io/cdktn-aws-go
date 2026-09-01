package awsbedrock


// Experimental.
type AwsBedrockEvaluationJob_RagConfigProperty struct {
	// knowledge_base_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#knowledge_base_config AwsBedrockEvaluationJob#knowledge_base_config}
	// Experimental.
	KnowledgeBaseConfig interface{} `field:"optional" json:"knowledgeBaseConfig" yaml:"knowledgeBaseConfig"`
	// precomputed_rag_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#precomputed_rag_source_config AwsBedrockEvaluationJob#precomputed_rag_source_config}
	// Experimental.
	PrecomputedRagSourceConfig interface{} `field:"optional" json:"precomputedRagSourceConfig" yaml:"precomputedRagSourceConfig"`
}

