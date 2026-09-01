package awstransferfamily


// Experimental.
type AwsTransferWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationEfsFileLocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#file_system_id AwsTransferWorkflow#file_system_id}.
	// Experimental.
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#path AwsTransferWorkflow#path}.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

