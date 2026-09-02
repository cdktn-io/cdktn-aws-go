package awssagemakerai


// Experimental.
type TfTrainingJob_ServerlessJobConfigProperty struct {
	// Base model ARN in SageMaker Public Hub. SageMaker always selects the latest version of the provided model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#base_model_arn TfTrainingJob#base_model_arn}
	// Experimental.
	BaseModelArn *string `field:"required" json:"baseModelArn" yaml:"baseModelArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#job_type TfTrainingJob#job_type}.
	// Experimental.
	JobType *string `field:"required" json:"jobType" yaml:"jobType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#accept_eula TfTrainingJob#accept_eula}.
	// Experimental.
	AcceptEula interface{} `field:"optional" json:"acceptEula" yaml:"acceptEula"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#customization_technique TfTrainingJob#customization_technique}.
	// Experimental.
	CustomizationTechnique *string `field:"optional" json:"customizationTechnique" yaml:"customizationTechnique"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#evaluation_type TfTrainingJob#evaluation_type}.
	// Experimental.
	EvaluationType *string `field:"optional" json:"evaluationType" yaml:"evaluationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#evaluator_arn TfTrainingJob#evaluator_arn}.
	// Experimental.
	EvaluatorArn *string `field:"optional" json:"evaluatorArn" yaml:"evaluatorArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#peft TfTrainingJob#peft}.
	// Experimental.
	Peft *string `field:"optional" json:"peft" yaml:"peft"`
}

