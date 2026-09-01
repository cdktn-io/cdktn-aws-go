package awssagemakerai


// Experimental.
type AwsSagemakerTrainingJob_MlflowConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#mlflow_resource_arn AwsSagemakerTrainingJob#mlflow_resource_arn}.
	// Experimental.
	MlflowResourceArn *string `field:"required" json:"mlflowResourceArn" yaml:"mlflowResourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#mlflow_experiment_name AwsSagemakerTrainingJob#mlflow_experiment_name}.
	// Experimental.
	MlflowExperimentName *string `field:"optional" json:"mlflowExperimentName" yaml:"mlflowExperimentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#mlflow_run_name AwsSagemakerTrainingJob#mlflow_run_name}.
	// Experimental.
	MlflowRunName *string `field:"optional" json:"mlflowRunName" yaml:"mlflowRunName"`
}

