package awsbedrockagentcore


// Experimental.
type TfEvaluator_BedrockEvaluatorModelConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#model_id TfEvaluator#model_id}.
	// Experimental.
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#additional_model_request_fields TfEvaluator#additional_model_request_fields}.
	// Experimental.
	AdditionalModelRequestFields *string `field:"optional" json:"additionalModelRequestFields" yaml:"additionalModelRequestFields"`
	// inference_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#inference_config TfEvaluator#inference_config}
	// Experimental.
	InferenceConfig interface{} `field:"optional" json:"inferenceConfig" yaml:"inferenceConfig"`
}

