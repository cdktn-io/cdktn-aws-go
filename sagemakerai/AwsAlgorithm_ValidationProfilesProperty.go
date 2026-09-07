package sagemakerai


// Experimental.
type AwsAlgorithm_ValidationProfilesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#profile_name AwsAlgorithm#profile_name}.
	// Experimental.
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	// training_job_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#training_job_definition AwsAlgorithm#training_job_definition}
	// Experimental.
	TrainingJobDefinition interface{} `field:"optional" json:"trainingJobDefinition" yaml:"trainingJobDefinition"`
	// transform_job_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#transform_job_definition AwsAlgorithm#transform_job_definition}
	// Experimental.
	TransformJobDefinition interface{} `field:"optional" json:"transformJobDefinition" yaml:"transformJobDefinition"`
}

