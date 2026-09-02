package awssigner


// Experimental.
type TfSigningJob_SourceS3Property struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/signer_signing_job#bucket TfSigningJob#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/signer_signing_job#key TfSigningJob#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/signer_signing_job#version TfSigningJob#version}.
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
}

