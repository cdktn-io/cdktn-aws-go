package awsbedrockagents


// Experimental.
type TfPrompt_ContentProperty struct {
	// cache_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#cache_point TfPrompt#cache_point}
	// Experimental.
	CachePoint interface{} `field:"optional" json:"cachePoint" yaml:"cachePoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#text TfPrompt#text}.
	// Experimental.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

