package awstransferfamily


// Experimental.
type TfWorkflow_OnExceptionStepsCustomStepDetailsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#name TfWorkflow#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#source_file_location TfWorkflow#source_file_location}.
	// Experimental.
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#target TfWorkflow#target}.
	// Experimental.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#timeout_seconds TfWorkflow#timeout_seconds}.
	// Experimental.
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

