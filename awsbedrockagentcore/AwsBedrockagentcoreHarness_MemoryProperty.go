package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreHarness_MemoryProperty struct {
	// agentcore_memory_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#agentcore_memory_configuration AwsBedrockagentcoreHarness#agentcore_memory_configuration}
	// Experimental.
	AgentcoreMemoryConfiguration interface{} `field:"optional" json:"agentcoreMemoryConfiguration" yaml:"agentcoreMemoryConfiguration"`
	// disabled block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#disabled AwsBedrockagentcoreHarness#disabled}
	// Experimental.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// managed_memory_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#managed_memory_configuration AwsBedrockagentcoreHarness#managed_memory_configuration}
	// Experimental.
	ManagedMemoryConfiguration interface{} `field:"optional" json:"managedMemoryConfiguration" yaml:"managedMemoryConfiguration"`
}

