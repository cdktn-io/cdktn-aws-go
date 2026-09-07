package bedrockagents


// Experimental.
type AwsAgent_InferenceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#max_length AwsAgent#max_length}.
	// Experimental.
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#stop_sequences AwsAgent#stop_sequences}.
	// Experimental.
	StopSequences *[]*string `field:"optional" json:"stopSequences" yaml:"stopSequences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#temperature AwsAgent#temperature}.
	// Experimental.
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#top_k AwsAgent#top_k}.
	// Experimental.
	TopK *float64 `field:"optional" json:"topK" yaml:"topK"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#top_p AwsAgent#top_p}.
	// Experimental.
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}

