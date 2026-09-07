package transferfamily


// Experimental.
type AwsWorkflow_StepsDecryptStepDetailsDestinationFileLocationProperty struct {
	// efs_file_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#efs_file_location AwsWorkflow#efs_file_location}
	// Experimental.
	EfsFileLocation *AwsWorkflow_StepsDecryptStepDetailsDestinationFileLocationEfsFileLocationProperty `field:"optional" json:"efsFileLocation" yaml:"efsFileLocation"`
	// s3_file_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#s3_file_location AwsWorkflow#s3_file_location}
	// Experimental.
	S3FileLocation *AwsWorkflow_StepsDecryptStepDetailsDestinationFileLocationS3FileLocationProperty `field:"optional" json:"s3FileLocation" yaml:"s3FileLocation"`
}

