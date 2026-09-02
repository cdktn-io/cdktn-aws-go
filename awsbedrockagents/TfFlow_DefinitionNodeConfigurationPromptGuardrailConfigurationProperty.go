package awsbedrockagents


// Experimental.
type TfFlow_DefinitionNodeConfigurationPromptGuardrailConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#guardrail_identifier TfFlow#guardrail_identifier}.
	// Experimental.
	GuardrailIdentifier *string `field:"required" json:"guardrailIdentifier" yaml:"guardrailIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#guardrail_version TfFlow#guardrail_version}.
	// Experimental.
	GuardrailVersion *string `field:"required" json:"guardrailVersion" yaml:"guardrailVersion"`
}

