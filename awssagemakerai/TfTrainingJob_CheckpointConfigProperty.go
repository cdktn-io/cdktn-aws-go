package awssagemakerai


// Experimental.
type TfTrainingJob_CheckpointConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#s3_uri TfTrainingJob#s3_uri}.
	// Experimental.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job#local_path TfTrainingJob#local_path}.
	// Experimental.
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
}

