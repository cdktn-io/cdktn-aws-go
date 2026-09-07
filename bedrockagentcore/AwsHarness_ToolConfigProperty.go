package bedrockagentcore


// Experimental.
type AwsHarness_ToolConfigProperty struct {
	// agentcore_browser block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#agentcore_browser AwsHarness#agentcore_browser}
	// Experimental.
	AgentcoreBrowser interface{} `field:"optional" json:"agentcoreBrowser" yaml:"agentcoreBrowser"`
	// agentcore_code_interpreter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#agentcore_code_interpreter AwsHarness#agentcore_code_interpreter}
	// Experimental.
	AgentcoreCodeInterpreter interface{} `field:"optional" json:"agentcoreCodeInterpreter" yaml:"agentcoreCodeInterpreter"`
	// agentcore_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#agentcore_gateway AwsHarness#agentcore_gateway}
	// Experimental.
	AgentcoreGateway interface{} `field:"optional" json:"agentcoreGateway" yaml:"agentcoreGateway"`
	// inline_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#inline_function AwsHarness#inline_function}
	// Experimental.
	InlineFunction interface{} `field:"optional" json:"inlineFunction" yaml:"inlineFunction"`
	// remote_mcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#remote_mcp AwsHarness#remote_mcp}
	// Experimental.
	RemoteMcp interface{} `field:"optional" json:"remoteMcp" yaml:"remoteMcp"`
}

