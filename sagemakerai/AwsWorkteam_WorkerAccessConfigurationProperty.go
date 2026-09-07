package sagemakerai


// Experimental.
type AwsWorkteam_WorkerAccessConfigurationProperty struct {
	// s3_presign block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workteam#s3_presign AwsWorkteam#s3_presign}
	// Experimental.
	S3Presign *AwsWorkteam_S3PresignProperty `field:"optional" json:"s3Presign" yaml:"s3Presign"`
}

