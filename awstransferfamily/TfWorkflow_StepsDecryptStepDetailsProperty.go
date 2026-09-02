package awstransferfamily


// Experimental.
type TfWorkflow_StepsDecryptStepDetailsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#type TfWorkflow#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// destination_file_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#destination_file_location TfWorkflow#destination_file_location}
	// Experimental.
	DestinationFileLocation *TfWorkflow_StepsDecryptStepDetailsDestinationFileLocationProperty `field:"optional" json:"destinationFileLocation" yaml:"destinationFileLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#name TfWorkflow#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#overwrite_existing TfWorkflow#overwrite_existing}.
	// Experimental.
	OverwriteExisting *string `field:"optional" json:"overwriteExisting" yaml:"overwriteExisting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#source_file_location TfWorkflow#source_file_location}.
	// Experimental.
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
}

