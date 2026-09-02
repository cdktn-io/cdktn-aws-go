package awsbedrockagents


// Experimental.
type TfFlow_InlineProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#model_id TfFlow#model_id}.
	// Experimental.
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#template_type TfFlow#template_type}.
	// Experimental.
	TemplateType *string `field:"required" json:"templateType" yaml:"templateType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#additional_model_request_fields TfFlow#additional_model_request_fields}.
	// Experimental.
	AdditionalModelRequestFields *string `field:"optional" json:"additionalModelRequestFields" yaml:"additionalModelRequestFields"`
	// inference_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#inference_configuration TfFlow#inference_configuration}
	// Experimental.
	InferenceConfiguration interface{} `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// template_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_flow#template_configuration TfFlow#template_configuration}
	// Experimental.
	TemplateConfiguration interface{} `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
}

