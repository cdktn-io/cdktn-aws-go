package awsbedrock


// Experimental.
type AwsBedrockEvaluationJob_AutomatedProperty struct {
	// custom_metric_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#custom_metric_config AwsBedrockEvaluationJob#custom_metric_config}
	// Experimental.
	CustomMetricConfig interface{} `field:"optional" json:"customMetricConfig" yaml:"customMetricConfig"`
	// dataset_metric_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#dataset_metric_config AwsBedrockEvaluationJob#dataset_metric_config}
	// Experimental.
	DatasetMetricConfig interface{} `field:"optional" json:"datasetMetricConfig" yaml:"datasetMetricConfig"`
	// evaluator_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#evaluator_model_config AwsBedrockEvaluationJob#evaluator_model_config}
	// Experimental.
	EvaluatorModelConfig interface{} `field:"optional" json:"evaluatorModelConfig" yaml:"evaluatorModelConfig"`
}

