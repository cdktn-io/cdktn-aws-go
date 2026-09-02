package awsbedrockagentcore


// Experimental.
type TfHarness_EnvironmentAgentcoreRuntimeEnvironmentLifecycleConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#idle_runtime_session_timeout TfHarness#idle_runtime_session_timeout}.
	// Experimental.
	IdleRuntimeSessionTimeout *float64 `field:"optional" json:"idleRuntimeSessionTimeout" yaml:"idleRuntimeSessionTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#max_lifetime TfHarness#max_lifetime}.
	// Experimental.
	MaxLifetime *float64 `field:"optional" json:"maxLifetime" yaml:"maxLifetime"`
}

