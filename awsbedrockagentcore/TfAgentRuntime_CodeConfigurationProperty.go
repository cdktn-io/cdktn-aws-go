package awsbedrockagentcore


// Experimental.
type TfAgentRuntime_CodeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#entry_point TfAgentRuntime#entry_point}.
	// Experimental.
	EntryPoint *[]*string `field:"required" json:"entryPoint" yaml:"entryPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#runtime TfAgentRuntime#runtime}.
	// Experimental.
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#code TfAgentRuntime#code}
	// Experimental.
	Code interface{} `field:"optional" json:"code" yaml:"code"`
}

