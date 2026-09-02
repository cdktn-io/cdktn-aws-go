package awsbedrockagents


// Experimental.
type TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#max_tokens TfFlow#max_tokens}.
	// Experimental.
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#stop_sequences TfFlow#stop_sequences}.
	// Experimental.
	StopSequences *[]*string `field:"optional" json:"stopSequences" yaml:"stopSequences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#temperature TfFlow#temperature}.
	// Experimental.
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#top_p TfFlow#top_p}.
	// Experimental.
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}

