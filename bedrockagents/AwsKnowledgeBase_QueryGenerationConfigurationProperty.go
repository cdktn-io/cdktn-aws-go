package bedrockagents


// Experimental.
type AwsKnowledgeBase_QueryGenerationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#execution_timeout_seconds AwsKnowledgeBase#execution_timeout_seconds}.
	// Experimental.
	ExecutionTimeoutSeconds *float64 `field:"optional" json:"executionTimeoutSeconds" yaml:"executionTimeoutSeconds"`
	// generation_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#generation_context AwsKnowledgeBase#generation_context}
	// Experimental.
	GenerationContext interface{} `field:"optional" json:"generationContext" yaml:"generationContext"`
}

