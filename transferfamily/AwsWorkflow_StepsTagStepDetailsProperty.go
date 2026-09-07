package transferfamily


// Experimental.
type AwsWorkflow_StepsTagStepDetailsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#name AwsWorkflow#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#source_file_location AwsWorkflow#source_file_location}.
	// Experimental.
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_workflow#tags AwsWorkflow#tags}
	// Experimental.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

