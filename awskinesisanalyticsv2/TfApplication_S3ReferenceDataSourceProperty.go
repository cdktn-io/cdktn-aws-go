package awskinesisanalyticsv2


// Experimental.
type TfApplication_S3ReferenceDataSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#bucket_arn TfApplication#bucket_arn}.
	// Experimental.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#file_key TfApplication#file_key}.
	// Experimental.
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
}

