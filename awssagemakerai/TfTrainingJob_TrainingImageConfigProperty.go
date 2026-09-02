package awssagemakerai


// Experimental.
type TfTrainingJob_TrainingImageConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#training_repository_access_mode TfTrainingJob#training_repository_access_mode}.
	// Experimental.
	TrainingRepositoryAccessMode *string `field:"optional" json:"trainingRepositoryAccessMode" yaml:"trainingRepositoryAccessMode"`
	// training_repository_auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#training_repository_auth_config TfTrainingJob#training_repository_auth_config}
	// Experimental.
	TrainingRepositoryAuthConfig interface{} `field:"optional" json:"trainingRepositoryAuthConfig" yaml:"trainingRepositoryAuthConfig"`
}

