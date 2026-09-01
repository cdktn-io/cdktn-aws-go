package awss3tables


// Experimental.
type AwsS3TablesTableBucket_MaintenanceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table_bucket#iceberg_unreferenced_file_removal AwsS3TablesTableBucket#iceberg_unreferenced_file_removal}.
	// Experimental.
	IcebergUnreferencedFileRemoval *AwsS3TablesTableBucket_IcebergUnreferencedFileRemovalProperty `field:"optional" json:"icebergUnreferencedFileRemoval" yaml:"icebergUnreferencedFileRemoval"`
}

