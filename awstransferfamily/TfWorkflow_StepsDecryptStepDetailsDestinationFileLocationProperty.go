package awstransferfamily


// Experimental.
type TfWorkflow_StepsDecryptStepDetailsDestinationFileLocationProperty struct {
	// efs_file_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#efs_file_location TfWorkflow#efs_file_location}
	// Experimental.
	EfsFileLocation *TfWorkflow_StepsDecryptStepDetailsDestinationFileLocationEfsFileLocationProperty `field:"optional" json:"efsFileLocation" yaml:"efsFileLocation"`
	// s3_file_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#s3_file_location TfWorkflow#s3_file_location}
	// Experimental.
	S3FileLocation *TfWorkflow_StepsDecryptStepDetailsDestinationFileLocationS3FileLocationProperty `field:"optional" json:"s3FileLocation" yaml:"s3FileLocation"`
}

