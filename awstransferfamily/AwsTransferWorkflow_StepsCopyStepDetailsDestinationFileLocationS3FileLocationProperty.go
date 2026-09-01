package awstransferfamily


// Experimental.
type AwsTransferWorkflow_StepsCopyStepDetailsDestinationFileLocationS3FileLocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#bucket AwsTransferWorkflow#bucket}.
	// Experimental.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#key AwsTransferWorkflow#key}.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

