package awsbedrockagents


// Experimental.
type AwsBedrockagentPrompt_TemplateConfigurationProperty struct {
	// chat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#chat AwsBedrockagentPrompt#chat}
	// Experimental.
	Chat interface{} `field:"optional" json:"chat" yaml:"chat"`
	// text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#text AwsBedrockagentPrompt#text}
	// Experimental.
	Text interface{} `field:"optional" json:"text" yaml:"text"`
}

