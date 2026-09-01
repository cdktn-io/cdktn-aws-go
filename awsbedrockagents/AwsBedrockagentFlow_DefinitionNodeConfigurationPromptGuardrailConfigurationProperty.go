package awsbedrockagents


// Experimental.
type AwsBedrockagentFlow_DefinitionNodeConfigurationPromptGuardrailConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#guardrail_identifier AwsBedrockagentFlow#guardrail_identifier}.
	// Experimental.
	GuardrailIdentifier *string `field:"required" json:"guardrailIdentifier" yaml:"guardrailIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#guardrail_version AwsBedrockagentFlow#guardrail_version}.
	// Experimental.
	GuardrailVersion *string `field:"required" json:"guardrailVersion" yaml:"guardrailVersion"`
}

