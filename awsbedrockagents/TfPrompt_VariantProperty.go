package awsbedrockagents


// Experimental.
type TfPrompt_VariantProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#name TfPrompt#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#template_type TfPrompt#template_type}.
	// Experimental.
	TemplateType *string `field:"required" json:"templateType" yaml:"templateType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#additional_model_request_fields TfPrompt#additional_model_request_fields}.
	// Experimental.
	AdditionalModelRequestFields *string `field:"optional" json:"additionalModelRequestFields" yaml:"additionalModelRequestFields"`
	// gen_ai_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#gen_ai_resource TfPrompt#gen_ai_resource}
	// Experimental.
	GenAiResource interface{} `field:"optional" json:"genAiResource" yaml:"genAiResource"`
	// inference_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#inference_configuration TfPrompt#inference_configuration}
	// Experimental.
	InferenceConfiguration interface{} `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#metadata TfPrompt#metadata}
	// Experimental.
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#model_id TfPrompt#model_id}.
	// Experimental.
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
	// template_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_prompt#template_configuration TfPrompt#template_configuration}
	// Experimental.
	TemplateConfiguration interface{} `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
}

