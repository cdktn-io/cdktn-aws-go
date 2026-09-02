package awsbedrockagents


// Experimental.
type TfAgent_InferenceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#max_length TfAgent#max_length}.
	// Experimental.
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#stop_sequences TfAgent#stop_sequences}.
	// Experimental.
	StopSequences *[]*string `field:"optional" json:"stopSequences" yaml:"stopSequences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#temperature TfAgent#temperature}.
	// Experimental.
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#top_k TfAgent#top_k}.
	// Experimental.
	TopK *float64 `field:"optional" json:"topK" yaml:"topK"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#top_p TfAgent#top_p}.
	// Experimental.
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}

