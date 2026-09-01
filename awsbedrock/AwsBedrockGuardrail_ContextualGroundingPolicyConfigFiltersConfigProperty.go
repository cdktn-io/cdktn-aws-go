package awsbedrock


// Experimental.
type AwsBedrockGuardrail_ContextualGroundingPolicyConfigFiltersConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#threshold AwsBedrockGuardrail#threshold}.
	// Experimental.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#type AwsBedrockGuardrail#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

