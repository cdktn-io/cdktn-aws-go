package awssagemakerai


// Experimental.
type TfTrainingJob_TensorBoardOutputConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#s3_output_path TfTrainingJob#s3_output_path}.
	// Experimental.
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#local_path TfTrainingJob#local_path}.
	// Experimental.
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
}

