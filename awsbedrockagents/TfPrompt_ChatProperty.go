package awsbedrockagents


// Experimental.
type TfPrompt_ChatProperty struct {
	// input_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#input_variable TfPrompt#input_variable}
	// Experimental.
	InputVariable interface{} `field:"optional" json:"inputVariable" yaml:"inputVariable"`
	// message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#message TfPrompt#message}
	// Experimental.
	Message interface{} `field:"optional" json:"message" yaml:"message"`
	// system block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#system TfPrompt#system}
	// Experimental.
	SystemAttribute interface{} `field:"optional" json:"systemAttribute" yaml:"systemAttribute"`
	// tool_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#tool_configuration TfPrompt#tool_configuration}
	// Experimental.
	ToolConfiguration interface{} `field:"optional" json:"toolConfiguration" yaml:"toolConfiguration"`
}

