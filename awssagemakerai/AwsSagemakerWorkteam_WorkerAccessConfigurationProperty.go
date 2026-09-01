package awssagemakerai


// Experimental.
type AwsSagemakerWorkteam_WorkerAccessConfigurationProperty struct {
	// s3_presign block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workteam#s3_presign AwsSagemakerWorkteam#s3_presign}
	// Experimental.
	S3Presign *AwsSagemakerWorkteam_S3PresignProperty `field:"optional" json:"s3Presign" yaml:"s3Presign"`
}

