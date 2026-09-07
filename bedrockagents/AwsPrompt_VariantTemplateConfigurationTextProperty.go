package bedrockagents


// Experimental.
type AwsPrompt_VariantTemplateConfigurationTextProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#text AwsPrompt#text}.
	// Experimental.
	Text *string `field:"required" json:"text" yaml:"text"`
	// cache_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#cache_point AwsPrompt#cache_point}
	// Experimental.
	CachePoint interface{} `field:"optional" json:"cachePoint" yaml:"cachePoint"`
	// input_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#input_variable AwsPrompt#input_variable}
	// Experimental.
	InputVariable interface{} `field:"optional" json:"inputVariable" yaml:"inputVariable"`
}

