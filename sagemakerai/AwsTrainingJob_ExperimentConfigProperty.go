package sagemakerai


// Experimental.
type AwsTrainingJob_ExperimentConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#experiment_name AwsTrainingJob#experiment_name}.
	// Experimental.
	ExperimentName *string `field:"optional" json:"experimentName" yaml:"experimentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#run_name AwsTrainingJob#run_name}.
	// Experimental.
	RunName *string `field:"optional" json:"runName" yaml:"runName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#trial_component_display_name AwsTrainingJob#trial_component_display_name}.
	// Experimental.
	TrialComponentDisplayName *string `field:"optional" json:"trialComponentDisplayName" yaml:"trialComponentDisplayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#trial_name AwsTrainingJob#trial_name}.
	// Experimental.
	TrialName *string `field:"optional" json:"trialName" yaml:"trialName"`
}

