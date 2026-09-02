package awsbedrockagents


// Experimental.
type TfPrompt_VariantTemplateConfigurationChatToolConfigurationToolProperty struct {
	// cache_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#cache_point TfPrompt#cache_point}
	// Experimental.
	CachePoint interface{} `field:"optional" json:"cachePoint" yaml:"cachePoint"`
	// tool_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#tool_spec TfPrompt#tool_spec}
	// Experimental.
	ToolSpec interface{} `field:"optional" json:"toolSpec" yaml:"toolSpec"`
}

