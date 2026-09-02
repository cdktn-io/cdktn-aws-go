package awssagemakerai


// Experimental.
type TfTrainingJob_MlflowConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#mlflow_resource_arn TfTrainingJob#mlflow_resource_arn}.
	// Experimental.
	MlflowResourceArn *string `field:"required" json:"mlflowResourceArn" yaml:"mlflowResourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#mlflow_experiment_name TfTrainingJob#mlflow_experiment_name}.
	// Experimental.
	MlflowExperimentName *string `field:"optional" json:"mlflowExperimentName" yaml:"mlflowExperimentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#mlflow_run_name TfTrainingJob#mlflow_run_name}.
	// Experimental.
	MlflowRunName *string `field:"optional" json:"mlflowRunName" yaml:"mlflowRunName"`
}

