package awsbedrockagents


// Experimental.
type TfFlow_ChatProperty struct {
	// input_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#input_variable TfFlow#input_variable}
	// Experimental.
	InputVariable interface{} `field:"optional" json:"inputVariable" yaml:"inputVariable"`
	// message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#message TfFlow#message}
	// Experimental.
	Message interface{} `field:"optional" json:"message" yaml:"message"`
	// system block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#system TfFlow#system}
	// Experimental.
	SystemAttribute interface{} `field:"optional" json:"systemAttribute" yaml:"systemAttribute"`
	// tool_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#tool_configuration TfFlow#tool_configuration}
	// Experimental.
	ToolConfiguration interface{} `field:"optional" json:"toolConfiguration" yaml:"toolConfiguration"`
}

