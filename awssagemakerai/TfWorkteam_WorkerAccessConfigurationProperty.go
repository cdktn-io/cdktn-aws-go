package awssagemakerai


// Experimental.
type TfWorkteam_WorkerAccessConfigurationProperty struct {
	// s3_presign block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workteam#s3_presign TfWorkteam#s3_presign}
	// Experimental.
	S3Presign *TfWorkteam_S3PresignProperty `field:"optional" json:"s3Presign" yaml:"s3Presign"`
}

