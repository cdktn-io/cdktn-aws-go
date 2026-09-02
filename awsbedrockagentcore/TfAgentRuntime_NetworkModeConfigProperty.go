package awsbedrockagentcore


// Experimental.
type TfAgentRuntime_NetworkModeConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#security_groups TfAgentRuntime#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime#subnets TfAgentRuntime#subnets}.
	// Experimental.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

