package awss3


// Experimental.
type TfBucketAnalyticsConfiguration_StorageClassAnalysisProperty struct {
	// data_export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_analytics_configuration#data_export TfBucketAnalyticsConfiguration#data_export}
	// Experimental.
	DataExport *TfBucketAnalyticsConfiguration_DataExportProperty `field:"required" json:"dataExport" yaml:"dataExport"`
}

