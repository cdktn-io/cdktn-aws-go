package awssagemakerai


// Experimental.
type TfLabelingJob_LabelingJobAlgorithmsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#labeling_job_algorithm_specification_arn TfLabelingJob#labeling_job_algorithm_specification_arn}.
	// Experimental.
	LabelingJobAlgorithmSpecificationArn *string `field:"required" json:"labelingJobAlgorithmSpecificationArn" yaml:"labelingJobAlgorithmSpecificationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#initial_active_learning_model_arn TfLabelingJob#initial_active_learning_model_arn}.
	// Experimental.
	InitialActiveLearningModelArn *string `field:"optional" json:"initialActiveLearningModelArn" yaml:"initialActiveLearningModelArn"`
	// labeling_job_resource_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#labeling_job_resource_config TfLabelingJob#labeling_job_resource_config}
	// Experimental.
	LabelingJobResourceConfig interface{} `field:"optional" json:"labelingJobResourceConfig" yaml:"labelingJobResourceConfig"`
}

