package sagemakerai


// Experimental.
type AwsTrainingJob_DebugHookConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#s3_output_path AwsTrainingJob#s3_output_path}.
	// Experimental.
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
	// collection_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#collection_configurations AwsTrainingJob#collection_configurations}
	// Experimental.
	CollectionConfigurations interface{} `field:"optional" json:"collectionConfigurations" yaml:"collectionConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#hook_parameters AwsTrainingJob#hook_parameters}.
	// Experimental.
	HookParameters *map[string]*string `field:"optional" json:"hookParameters" yaml:"hookParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#local_path AwsTrainingJob#local_path}.
	// Experimental.
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
}

