package awsbedrock


// Experimental.
type TfGuardrail_TopicsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#definition TfGuardrail#definition}.
	// Experimental.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#name TfGuardrail#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#type TfGuardrail#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#examples TfGuardrail#examples}.
	// Experimental.
	Examples *[]*string `field:"optional" json:"examples" yaml:"examples"`
}

