package bedrockagents


// Experimental.
type AwsAgent_PromptConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#base_prompt_template AwsAgent#base_prompt_template}.
	// Experimental.
	BasePromptTemplate *string `field:"optional" json:"basePromptTemplate" yaml:"basePromptTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#inference_configuration AwsAgent#inference_configuration}.
	// Experimental.
	InferenceConfiguration interface{} `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#parser_mode AwsAgent#parser_mode}.
	// Experimental.
	ParserMode *string `field:"optional" json:"parserMode" yaml:"parserMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#prompt_creation_mode AwsAgent#prompt_creation_mode}.
	// Experimental.
	PromptCreationMode *string `field:"optional" json:"promptCreationMode" yaml:"promptCreationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#prompt_state AwsAgent#prompt_state}.
	// Experimental.
	PromptState *string `field:"optional" json:"promptState" yaml:"promptState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#prompt_type AwsAgent#prompt_type}.
	// Experimental.
	PromptType *string `field:"optional" json:"promptType" yaml:"promptType"`
}

