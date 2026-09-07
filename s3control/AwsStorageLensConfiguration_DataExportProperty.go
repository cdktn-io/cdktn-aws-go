package s3control


// Experimental.
type AwsStorageLensConfiguration_DataExportProperty struct {
	// cloud_watch_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#cloud_watch_metrics AwsStorageLensConfiguration#cloud_watch_metrics}
	// Experimental.
	CloudWatchMetrics *AwsStorageLensConfiguration_CloudWatchMetricsProperty `field:"optional" json:"cloudWatchMetrics" yaml:"cloudWatchMetrics"`
	// s3_bucket_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#s3_bucket_destination AwsStorageLensConfiguration#s3_bucket_destination}
	// Experimental.
	S3BucketDestination *AwsStorageLensConfiguration_StorageLensConfigurationDataExportS3BucketDestinationProperty `field:"optional" json:"s3BucketDestination" yaml:"s3BucketDestination"`
	// storage_lens_table_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#storage_lens_table_destination AwsStorageLensConfiguration#storage_lens_table_destination}
	// Experimental.
	StorageLensTableDestination *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationProperty `field:"optional" json:"storageLensTableDestination" yaml:"storageLensTableDestination"`
}

