package awsbedrockagents


// Experimental.
type TfAgent_GuardrailConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#guardrail_identifier TfAgent#guardrail_identifier}.
	// Experimental.
	GuardrailIdentifier *string `field:"optional" json:"guardrailIdentifier" yaml:"guardrailIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#guardrail_version TfAgent#guardrail_version}.
	// Experimental.
	GuardrailVersion *string `field:"optional" json:"guardrailVersion" yaml:"guardrailVersion"`
}

