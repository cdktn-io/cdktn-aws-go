package bedrock


// Experimental.
type AwsEvaluationJob_EvaluationConfigHumanDatasetMetricConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#metric_names AwsEvaluationJob#metric_names}.
	// Experimental.
	MetricNames *[]*string `field:"required" json:"metricNames" yaml:"metricNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#task_type AwsEvaluationJob#task_type}.
	// Experimental.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// dataset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#dataset AwsEvaluationJob#dataset}
	// Experimental.
	Dataset interface{} `field:"optional" json:"dataset" yaml:"dataset"`
}

