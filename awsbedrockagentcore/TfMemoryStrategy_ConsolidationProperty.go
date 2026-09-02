package awsbedrockagentcore


// Experimental.
type TfMemoryStrategy_ConsolidationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#append_to_prompt TfMemoryStrategy#append_to_prompt}.
	// Experimental.
	AppendToPrompt *string `field:"required" json:"appendToPrompt" yaml:"appendToPrompt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#model_id TfMemoryStrategy#model_id}.
	// Experimental.
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
}

