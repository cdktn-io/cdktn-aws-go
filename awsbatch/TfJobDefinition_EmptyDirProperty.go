package awsbatch


// Experimental.
type TfJobDefinition_EmptyDirProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#size_limit TfJobDefinition#size_limit}.
	// Experimental.
	SizeLimit *string `field:"required" json:"sizeLimit" yaml:"sizeLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#medium TfJobDefinition#medium}.
	// Experimental.
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
}

