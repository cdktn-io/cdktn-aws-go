package bedrockagents


// Experimental.
type AwsFlow_DefinitionNodeConfigurationKnowledgeBaseGuardrailConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#guardrail_identifier AwsFlow#guardrail_identifier}.
	// Experimental.
	GuardrailIdentifier *string `field:"required" json:"guardrailIdentifier" yaml:"guardrailIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#guardrail_version AwsFlow#guardrail_version}.
	// Experimental.
	GuardrailVersion *string `field:"required" json:"guardrailVersion" yaml:"guardrailVersion"`
}

