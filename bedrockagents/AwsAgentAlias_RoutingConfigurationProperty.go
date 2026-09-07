package bedrockagents


// Experimental.
type AwsAgentAlias_RoutingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_alias#agent_version AwsAgentAlias#agent_version}.
	// Experimental.
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_alias#provisioned_throughput AwsAgentAlias#provisioned_throughput}.
	// Experimental.
	ProvisionedThroughput *string `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
}

