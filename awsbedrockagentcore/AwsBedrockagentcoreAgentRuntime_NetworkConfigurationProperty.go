package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreAgentRuntime_NetworkConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#network_mode AwsBedrockagentcoreAgentRuntime#network_mode}.
	// Experimental.
	NetworkMode *string `field:"required" json:"networkMode" yaml:"networkMode"`
	// network_mode_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#network_mode_config AwsBedrockagentcoreAgentRuntime#network_mode_config}
	// Experimental.
	NetworkModeConfig interface{} `field:"optional" json:"networkModeConfig" yaml:"networkModeConfig"`
}

