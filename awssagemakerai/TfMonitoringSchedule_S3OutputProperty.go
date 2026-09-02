package awssagemakerai


// Experimental.
type TfMonitoringSchedule_S3OutputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#local_path TfMonitoringSchedule#local_path}.
	// Experimental.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#s3_uri TfMonitoringSchedule#s3_uri}.
	// Experimental.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_monitoring_schedule#s3_upload_mode TfMonitoringSchedule#s3_upload_mode}.
	// Experimental.
	S3UploadMode *string `field:"optional" json:"s3UploadMode" yaml:"s3UploadMode"`
}

