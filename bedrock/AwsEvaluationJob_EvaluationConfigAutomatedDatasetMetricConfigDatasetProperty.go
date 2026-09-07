package bedrock


// Experimental.
type AwsEvaluationJob_EvaluationConfigAutomatedDatasetMetricConfigDatasetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#name AwsEvaluationJob#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// dataset_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#dataset_location AwsEvaluationJob#dataset_location}
	// Experimental.
	DatasetLocation interface{} `field:"optional" json:"datasetLocation" yaml:"datasetLocation"`
}

