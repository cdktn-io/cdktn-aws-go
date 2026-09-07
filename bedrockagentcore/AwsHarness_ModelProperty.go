package bedrockagentcore


// Experimental.
type AwsHarness_ModelProperty struct {
	// bedrock_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#bedrock_model_config AwsHarness#bedrock_model_config}
	// Experimental.
	BedrockModelConfig interface{} `field:"optional" json:"bedrockModelConfig" yaml:"bedrockModelConfig"`
	// gemini_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#gemini_model_config AwsHarness#gemini_model_config}
	// Experimental.
	GeminiModelConfig interface{} `field:"optional" json:"geminiModelConfig" yaml:"geminiModelConfig"`
	// openai_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#openai_model_config AwsHarness#openai_model_config}
	// Experimental.
	OpenaiModelConfig interface{} `field:"optional" json:"openaiModelConfig" yaml:"openaiModelConfig"`
}

