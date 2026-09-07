package bedrockagents


// Experimental.
type AwsAgentActionGroup_ActionGroupExecutorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#custom_control AwsAgentActionGroup#custom_control}.
	// Experimental.
	CustomControl *string `field:"optional" json:"customControl" yaml:"customControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#lambda AwsAgentActionGroup#lambda}.
	// Experimental.
	Lambda *string `field:"optional" json:"lambda" yaml:"lambda"`
}

