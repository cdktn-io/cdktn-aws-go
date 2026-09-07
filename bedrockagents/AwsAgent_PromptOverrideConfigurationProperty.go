package bedrockagents


// Experimental.
type AwsAgent_PromptOverrideConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#override_lambda AwsAgent#override_lambda}.
	// Experimental.
	OverrideLambda *string `field:"optional" json:"overrideLambda" yaml:"overrideLambda"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#prompt_configurations AwsAgent#prompt_configurations}.
	// Experimental.
	PromptConfigurations interface{} `field:"optional" json:"promptConfigurations" yaml:"promptConfigurations"`
}

