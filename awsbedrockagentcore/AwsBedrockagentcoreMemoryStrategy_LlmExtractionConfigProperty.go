package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreMemoryStrategy_LlmExtractionConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#definition AwsBedrockagentcoreMemoryStrategy#definition}.
	// Experimental.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#llm_extraction_instruction AwsBedrockagentcoreMemoryStrategy#llm_extraction_instruction}.
	// Experimental.
	LlmExtractionInstruction *string `field:"optional" json:"llmExtractionInstruction" yaml:"llmExtractionInstruction"`
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#validation AwsBedrockagentcoreMemoryStrategy#validation}
	// Experimental.
	Validation interface{} `field:"optional" json:"validation" yaml:"validation"`
}

