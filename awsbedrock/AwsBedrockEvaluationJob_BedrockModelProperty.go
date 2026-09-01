package awsbedrock


// Experimental.
type AwsBedrockEvaluationJob_BedrockModelProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#model_identifier AwsBedrockEvaluationJob#model_identifier}.
	// Experimental.
	ModelIdentifier *string `field:"required" json:"modelIdentifier" yaml:"modelIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#inference_params AwsBedrockEvaluationJob#inference_params}.
	// Experimental.
	InferenceParams *string `field:"optional" json:"inferenceParams" yaml:"inferenceParams"`
	// performance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#performance_config AwsBedrockEvaluationJob#performance_config}
	// Experimental.
	PerformanceConfig interface{} `field:"optional" json:"performanceConfig" yaml:"performanceConfig"`
}

