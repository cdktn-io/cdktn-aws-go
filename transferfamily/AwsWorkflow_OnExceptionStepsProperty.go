package transferfamily


// Experimental.
type AwsWorkflow_OnExceptionStepsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#type AwsWorkflow#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// copy_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#copy_step_details AwsWorkflow#copy_step_details}
	// Experimental.
	CopyStepDetails *AwsWorkflow_OnExceptionStepsCopyStepDetailsProperty `field:"optional" json:"copyStepDetails" yaml:"copyStepDetails"`
	// custom_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#custom_step_details AwsWorkflow#custom_step_details}
	// Experimental.
	CustomStepDetails *AwsWorkflow_OnExceptionStepsCustomStepDetailsProperty `field:"optional" json:"customStepDetails" yaml:"customStepDetails"`
	// decrypt_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#decrypt_step_details AwsWorkflow#decrypt_step_details}
	// Experimental.
	DecryptStepDetails *AwsWorkflow_OnExceptionStepsDecryptStepDetailsProperty `field:"optional" json:"decryptStepDetails" yaml:"decryptStepDetails"`
	// delete_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#delete_step_details AwsWorkflow#delete_step_details}
	// Experimental.
	DeleteStepDetails *AwsWorkflow_OnExceptionStepsDeleteStepDetailsProperty `field:"optional" json:"deleteStepDetails" yaml:"deleteStepDetails"`
	// tag_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#tag_step_details AwsWorkflow#tag_step_details}
	// Experimental.
	TagStepDetails *AwsWorkflow_OnExceptionStepsTagStepDetailsProperty `field:"optional" json:"tagStepDetails" yaml:"tagStepDetails"`
}

