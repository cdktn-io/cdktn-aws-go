package awsbedrock


// Experimental.
type AwsBedrockEvaluationJob_HumanProperty struct {
	// custom_metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#custom_metric AwsBedrockEvaluationJob#custom_metric}
	// Experimental.
	CustomMetric interface{} `field:"optional" json:"customMetric" yaml:"customMetric"`
	// dataset_metric_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#dataset_metric_config AwsBedrockEvaluationJob#dataset_metric_config}
	// Experimental.
	DatasetMetricConfig interface{} `field:"optional" json:"datasetMetricConfig" yaml:"datasetMetricConfig"`
	// human_workflow_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#human_workflow_config AwsBedrockEvaluationJob#human_workflow_config}
	// Experimental.
	HumanWorkflowConfig interface{} `field:"optional" json:"humanWorkflowConfig" yaml:"humanWorkflowConfig"`
}

