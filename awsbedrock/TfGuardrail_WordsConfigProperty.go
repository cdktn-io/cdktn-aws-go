package awsbedrock


// Experimental.
type TfGuardrail_WordsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#text TfGuardrail#text}.
	// Experimental.
	Text *string `field:"required" json:"text" yaml:"text"`
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

