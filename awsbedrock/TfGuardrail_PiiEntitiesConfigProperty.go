package awsbedrock


// Experimental.
type TfGuardrail_PiiEntitiesConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#action TfGuardrail#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#type TfGuardrail#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#input_action TfGuardrail#input_action}.
	// Experimental.
	InputAction *string `field:"optional" json:"inputAction" yaml:"inputAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#input_enabled TfGuardrail#input_enabled}.
	// Experimental.
	InputEnabled interface{} `field:"optional" json:"inputEnabled" yaml:"inputEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#output_action TfGuardrail#output_action}.
	// Experimental.
	OutputAction *string `field:"optional" json:"outputAction" yaml:"outputAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#output_enabled TfGuardrail#output_enabled}.
	// Experimental.
	OutputEnabled interface{} `field:"optional" json:"outputEnabled" yaml:"outputEnabled"`
}

