package awsbedrockagents


// Experimental.
type TfAgent_PromptOverrideConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#override_lambda TfAgent#override_lambda}.
	// Experimental.
	OverrideLambda *string `field:"optional" json:"overrideLambda" yaml:"overrideLambda"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#prompt_configurations TfAgent#prompt_configurations}.
	// Experimental.
	PromptConfigurations interface{} `field:"optional" json:"promptConfigurations" yaml:"promptConfigurations"`
}

