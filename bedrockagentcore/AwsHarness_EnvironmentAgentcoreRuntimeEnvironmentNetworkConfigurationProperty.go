package bedrockagentcore


// Experimental.
type AwsHarness_EnvironmentAgentcoreRuntimeEnvironmentNetworkConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#network_mode AwsHarness#network_mode}.
	// Experimental.
	NetworkMode *string `field:"required" json:"networkMode" yaml:"networkMode"`
	// network_mode_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#network_mode_config AwsHarness#network_mode_config}
	// Experimental.
	NetworkModeConfig interface{} `field:"optional" json:"networkModeConfig" yaml:"networkModeConfig"`
}

