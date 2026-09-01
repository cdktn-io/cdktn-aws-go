package awss3


// Experimental.
type AwsS3BucketAnalyticsConfiguration_StorageClassAnalysisProperty struct {
	// data_export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_analytics_configuration#data_export AwsS3BucketAnalyticsConfiguration#data_export}
	// Experimental.
	DataExport *AwsS3BucketAnalyticsConfiguration_DataExportProperty `field:"required" json:"dataExport" yaml:"dataExport"`
}

