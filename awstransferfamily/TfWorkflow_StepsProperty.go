package awstransferfamily


// Experimental.
type TfWorkflow_StepsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#type TfWorkflow#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// copy_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#copy_step_details TfWorkflow#copy_step_details}
	// Experimental.
	CopyStepDetails *TfWorkflow_StepsCopyStepDetailsProperty `field:"optional" json:"copyStepDetails" yaml:"copyStepDetails"`
	// custom_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#custom_step_details TfWorkflow#custom_step_details}
	// Experimental.
	CustomStepDetails *TfWorkflow_StepsCustomStepDetailsProperty `field:"optional" json:"customStepDetails" yaml:"customStepDetails"`
	// decrypt_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#decrypt_step_details TfWorkflow#decrypt_step_details}
	// Experimental.
	DecryptStepDetails *TfWorkflow_StepsDecryptStepDetailsProperty `field:"optional" json:"decryptStepDetails" yaml:"decryptStepDetails"`
	// delete_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#delete_step_details TfWorkflow#delete_step_details}
	// Experimental.
	DeleteStepDetails *TfWorkflow_StepsDeleteStepDetailsProperty `field:"optional" json:"deleteStepDetails" yaml:"deleteStepDetails"`
	// tag_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#tag_step_details TfWorkflow#tag_step_details}
	// Experimental.
	TagStepDetails *TfWorkflow_StepsTagStepDetailsProperty `field:"optional" json:"tagStepDetails" yaml:"tagStepDetails"`
}

