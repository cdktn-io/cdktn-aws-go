package sagemakerai


// Experimental.
type AwsTrainingJob_MlflowConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#mlflow_resource_arn AwsTrainingJob#mlflow_resource_arn}.
	// Experimental.
	MlflowResourceArn *string `field:"required" json:"mlflowResourceArn" yaml:"mlflowResourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#mlflow_experiment_name AwsTrainingJob#mlflow_experiment_name}.
	// Experimental.
	MlflowExperimentName *string `field:"optional" json:"mlflowExperimentName" yaml:"mlflowExperimentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#mlflow_run_name AwsTrainingJob#mlflow_run_name}.
	// Experimental.
	MlflowRunName *string `field:"optional" json:"mlflowRunName" yaml:"mlflowRunName"`
}

