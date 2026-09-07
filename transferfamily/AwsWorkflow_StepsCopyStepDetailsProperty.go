package transferfamily


// Experimental.
type AwsWorkflow_StepsCopyStepDetailsProperty struct {
	// destination_file_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#destination_file_location AwsWorkflow#destination_file_location}
	// Experimental.
	DestinationFileLocation *AwsWorkflow_StepsCopyStepDetailsDestinationFileLocationProperty `field:"optional" json:"destinationFileLocation" yaml:"destinationFileLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#name AwsWorkflow#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#overwrite_existing AwsWorkflow#overwrite_existing}.
	// Experimental.
	OverwriteExisting *string `field:"optional" json:"overwriteExisting" yaml:"overwriteExisting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#source_file_location AwsWorkflow#source_file_location}.
	// Experimental.
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
}

