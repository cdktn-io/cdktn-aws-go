package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreHarness_ModelProperty struct {
	// bedrock_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#bedrock_model_config AwsBedrockagentcoreHarness#bedrock_model_config}
	// Experimental.
	BedrockModelConfig interface{} `field:"optional" json:"bedrockModelConfig" yaml:"bedrockModelConfig"`
	// gemini_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#gemini_model_config AwsBedrockagentcoreHarness#gemini_model_config}
	// Experimental.
	GeminiModelConfig interface{} `field:"optional" json:"geminiModelConfig" yaml:"geminiModelConfig"`
	// openai_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#openai_model_config AwsBedrockagentcoreHarness#openai_model_config}
	// Experimental.
	OpenaiModelConfig interface{} `field:"optional" json:"openaiModelConfig" yaml:"openaiModelConfig"`
}

