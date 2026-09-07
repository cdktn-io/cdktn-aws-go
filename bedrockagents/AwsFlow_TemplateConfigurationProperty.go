package bedrockagents


// Experimental.
type AwsFlow_TemplateConfigurationProperty struct {
	// chat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#chat AwsFlow#chat}
	// Experimental.
	Chat interface{} `field:"optional" json:"chat" yaml:"chat"`
	// text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#text AwsFlow#text}
	// Experimental.
	Text interface{} `field:"optional" json:"text" yaml:"text"`
}

