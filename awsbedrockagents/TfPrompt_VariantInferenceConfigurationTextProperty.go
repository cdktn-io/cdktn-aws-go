package awsbedrockagents


// Experimental.
type TfPrompt_VariantInferenceConfigurationTextProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#max_tokens TfPrompt#max_tokens}.
	// Experimental.
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#stop_sequences TfPrompt#stop_sequences}.
	// Experimental.
	StopSequences *[]*string `field:"optional" json:"stopSequences" yaml:"stopSequences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#temperature TfPrompt#temperature}.
	// Experimental.
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#top_p TfPrompt#top_p}.
	// Experimental.
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}

