package awsbedrockagentcore


// Experimental.
type TfEvaluator_InferenceConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#max_tokens TfEvaluator#max_tokens}.
	// Experimental.
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#stop_sequences TfEvaluator#stop_sequences}.
	// Experimental.
	StopSequences *[]*string `field:"optional" json:"stopSequences" yaml:"stopSequences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#temperature TfEvaluator#temperature}.
	// Experimental.
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#top_p TfEvaluator#top_p}.
	// Experimental.
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}

