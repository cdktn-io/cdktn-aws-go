package s3tables


// Experimental.
type AwsTableBucket_MaintenanceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table_bucket#iceberg_unreferenced_file_removal AwsTableBucket#iceberg_unreferenced_file_removal}.
	// Experimental.
	IcebergUnreferencedFileRemoval *AwsTableBucket_IcebergUnreferencedFileRemovalProperty `field:"optional" json:"icebergUnreferencedFileRemoval" yaml:"icebergUnreferencedFileRemoval"`
}

