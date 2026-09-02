package awss3tables


// Experimental.
type TfTableBucket_MaintenanceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table_bucket#iceberg_unreferenced_file_removal TfTableBucket#iceberg_unreferenced_file_removal}.
	// Experimental.
	IcebergUnreferencedFileRemoval *TfTableBucket_IcebergUnreferencedFileRemovalProperty `field:"optional" json:"icebergUnreferencedFileRemoval" yaml:"icebergUnreferencedFileRemoval"`
}

