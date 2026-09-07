package s3


// Experimental.
type AwsBucketAnalyticsConfiguration_StorageClassAnalysisProperty struct {
	// data_export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_analytics_configuration#data_export AwsBucketAnalyticsConfiguration#data_export}
	// Experimental.
	DataExport *AwsBucketAnalyticsConfiguration_DataExportProperty `field:"required" json:"dataExport" yaml:"dataExport"`
}

