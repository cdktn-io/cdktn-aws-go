package awstransferfamily


// Experimental.
type AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationProperty struct {
	// efs_file_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#efs_file_location AwsTransferWorkflow#efs_file_location}
	// Experimental.
	EfsFileLocation *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationProperty `field:"optional" json:"efsFileLocation" yaml:"efsFileLocation"`
	// s3_file_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#s3_file_location AwsTransferWorkflow#s3_file_location}
	// Experimental.
	S3FileLocation *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationS3FileLocationProperty `field:"optional" json:"s3FileLocation" yaml:"s3FileLocation"`
}

