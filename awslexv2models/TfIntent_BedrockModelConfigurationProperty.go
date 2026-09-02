package awslexv2models


// Experimental.
type TfIntent_BedrockModelConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#model_arn TfIntent#model_arn}.
	// Experimental.
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#custom_prompt TfIntent#custom_prompt}.
	// Experimental.
	CustomPrompt *string `field:"optional" json:"customPrompt" yaml:"customPrompt"`
	// guardrail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#guardrail TfIntent#guardrail}
	// Experimental.
	Guardrail interface{} `field:"optional" json:"guardrail" yaml:"guardrail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#trace_status TfIntent#trace_status}.
	// Experimental.
	TraceStatus *string `field:"optional" json:"traceStatus" yaml:"traceStatus"`
}

