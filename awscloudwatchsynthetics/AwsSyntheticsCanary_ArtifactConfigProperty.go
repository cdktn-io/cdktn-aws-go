package awscloudwatchsynthetics


// Experimental.
type AwsSyntheticsCanary_ArtifactConfigProperty struct {
	// s3_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary#s3_encryption AwsSyntheticsCanary#s3_encryption}
	// Experimental.
	S3Encryption *AwsSyntheticsCanary_S3EncryptionProperty `field:"optional" json:"s3Encryption" yaml:"s3Encryption"`
}

