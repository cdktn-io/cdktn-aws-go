package awsbedrockagents


// Experimental.
type AwsBedrockagentFlow_SystemProperty struct {
	// cache_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#cache_point AwsBedrockagentFlow#cache_point}
	// Experimental.
	CachePoint interface{} `field:"optional" json:"cachePoint" yaml:"cachePoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#text AwsBedrockagentFlow#text}.
	// Experimental.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

