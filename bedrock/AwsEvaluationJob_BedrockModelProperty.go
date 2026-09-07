package bedrock


// Experimental.
type AwsEvaluationJob_BedrockModelProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#model_identifier AwsEvaluationJob#model_identifier}.
	// Experimental.
	ModelIdentifier *string `field:"required" json:"modelIdentifier" yaml:"modelIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#inference_params AwsEvaluationJob#inference_params}.
	// Experimental.
	InferenceParams *string `field:"optional" json:"inferenceParams" yaml:"inferenceParams"`
	// performance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#performance_config AwsEvaluationJob#performance_config}
	// Experimental.
	PerformanceConfig interface{} `field:"optional" json:"performanceConfig" yaml:"performanceConfig"`
}

