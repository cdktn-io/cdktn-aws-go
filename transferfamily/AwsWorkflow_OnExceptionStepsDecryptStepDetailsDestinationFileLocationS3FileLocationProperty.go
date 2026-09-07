package transferfamily


// Experimental.
type AwsWorkflow_OnExceptionStepsDecryptStepDetailsDestinationFileLocationS3FileLocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#bucket AwsWorkflow#bucket}.
	// Experimental.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#key AwsWorkflow#key}.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

