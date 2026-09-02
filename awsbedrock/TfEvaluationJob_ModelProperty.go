package awsbedrock


// Experimental.
type TfEvaluationJob_ModelProperty struct {
	// bedrock_model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#bedrock_model TfEvaluationJob#bedrock_model}
	// Experimental.
	BedrockModel interface{} `field:"optional" json:"bedrockModel" yaml:"bedrockModel"`
	// precomputed_inference_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#precomputed_inference_source TfEvaluationJob#precomputed_inference_source}
	// Experimental.
	PrecomputedInferenceSource interface{} `field:"optional" json:"precomputedInferenceSource" yaml:"precomputedInferenceSource"`
}

