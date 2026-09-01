package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreAgentRuntime_LifecycleConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#idle_runtime_session_timeout AwsBedrockagentcoreAgentRuntime#idle_runtime_session_timeout}.
	// Experimental.
	IdleRuntimeSessionTimeout *float64 `field:"optional" json:"idleRuntimeSessionTimeout" yaml:"idleRuntimeSessionTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#max_lifetime AwsBedrockagentcoreAgentRuntime#max_lifetime}.
	// Experimental.
	MaxLifetime *float64 `field:"optional" json:"maxLifetime" yaml:"maxLifetime"`
}

