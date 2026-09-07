package bedrock


// Experimental.
type AwsGuardrail_ContextualGroundingPolicyConfigFiltersConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#threshold AwsGuardrail#threshold}.
	// Experimental.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#type AwsGuardrail#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

