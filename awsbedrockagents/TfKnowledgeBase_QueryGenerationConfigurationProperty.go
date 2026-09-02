package awsbedrockagents


// Experimental.
type TfKnowledgeBase_QueryGenerationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#execution_timeout_seconds TfKnowledgeBase#execution_timeout_seconds}.
	// Experimental.
	ExecutionTimeoutSeconds *float64 `field:"optional" json:"executionTimeoutSeconds" yaml:"executionTimeoutSeconds"`
	// generation_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#generation_context TfKnowledgeBase#generation_context}
	// Experimental.
	GenerationContext interface{} `field:"optional" json:"generationContext" yaml:"generationContext"`
}

