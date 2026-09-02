package awstransferfamily


// Experimental.
type TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationProperty struct {
	// efs_file_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#efs_file_location TfWorkflow#efs_file_location}
	// Experimental.
	EfsFileLocation *TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationEfsFileLocationProperty `field:"optional" json:"efsFileLocation" yaml:"efsFileLocation"`
	// s3_file_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#s3_file_location TfWorkflow#s3_file_location}
	// Experimental.
	S3FileLocation *TfWorkflow_OnExceptionStepsCopyStepDetailsDestinationFileLocationS3FileLocationProperty `field:"optional" json:"s3FileLocation" yaml:"s3FileLocation"`
}

