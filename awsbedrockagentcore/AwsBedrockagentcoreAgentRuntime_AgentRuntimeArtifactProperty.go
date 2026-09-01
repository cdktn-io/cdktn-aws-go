package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreAgentRuntime_AgentRuntimeArtifactProperty struct {
	// code_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#code_configuration AwsBedrockagentcoreAgentRuntime#code_configuration}
	// Experimental.
	CodeConfiguration interface{} `field:"optional" json:"codeConfiguration" yaml:"codeConfiguration"`
	// container_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#container_configuration AwsBedrockagentcoreAgentRuntime#container_configuration}
	// Experimental.
	ContainerConfiguration interface{} `field:"optional" json:"containerConfiguration" yaml:"containerConfiguration"`
}

