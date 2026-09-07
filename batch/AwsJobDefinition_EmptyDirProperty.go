package batch


// Experimental.
type AwsJobDefinition_EmptyDirProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#size_limit AwsJobDefinition#size_limit}.
	// Experimental.
	SizeLimit *string `field:"required" json:"sizeLimit" yaml:"sizeLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#medium AwsJobDefinition#medium}.
	// Experimental.
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
}

