package awstransferfamily


// Experimental.
type AwsTransferWorkflow_OnExceptionStepsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#type AwsTransferWorkflow#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// copy_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#copy_step_details AwsTransferWorkflow#copy_step_details}
	// Experimental.
	CopyStepDetails *AwsTransferWorkflow_OnExceptionStepsCopyStepDetailsProperty `field:"optional" json:"copyStepDetails" yaml:"copyStepDetails"`
	// custom_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#custom_step_details AwsTransferWorkflow#custom_step_details}
	// Experimental.
	CustomStepDetails *AwsTransferWorkflow_OnExceptionStepsCustomStepDetailsProperty `field:"optional" json:"customStepDetails" yaml:"customStepDetails"`
	// decrypt_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#decrypt_step_details AwsTransferWorkflow#decrypt_step_details}
	// Experimental.
	DecryptStepDetails *AwsTransferWorkflow_OnExceptionStepsDecryptStepDetailsProperty `field:"optional" json:"decryptStepDetails" yaml:"decryptStepDetails"`
	// delete_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#delete_step_details AwsTransferWorkflow#delete_step_details}
	// Experimental.
	DeleteStepDetails *AwsTransferWorkflow_OnExceptionStepsDeleteStepDetailsProperty `field:"optional" json:"deleteStepDetails" yaml:"deleteStepDetails"`
	// tag_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#tag_step_details AwsTransferWorkflow#tag_step_details}
	// Experimental.
	TagStepDetails *AwsTransferWorkflow_OnExceptionStepsTagStepDetailsProperty `field:"optional" json:"tagStepDetails" yaml:"tagStepDetails"`
}

