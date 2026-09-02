package awsbedrock


// Experimental.
type TfEvaluationJob_CustomMetricConfigProperty struct {
	// custom_metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#custom_metric TfEvaluationJob#custom_metric}
	// Experimental.
	CustomMetric interface{} `field:"optional" json:"customMetric" yaml:"customMetric"`
	// evaluator_model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#evaluator_model_config TfEvaluationJob#evaluator_model_config}
	// Experimental.
	EvaluatorModelConfig interface{} `field:"optional" json:"evaluatorModelConfig" yaml:"evaluatorModelConfig"`
}

