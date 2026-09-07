package signer


// Experimental.
type AwsSigningJob_DestinationProperty struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/signer_signing_job#s3 AwsSigningJob#s3}
	// Experimental.
	S3 *AwsSigningJob_DestinationS3Property `field:"required" json:"s3" yaml:"s3"`
}

