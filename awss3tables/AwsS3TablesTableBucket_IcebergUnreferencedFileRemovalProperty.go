package awss3tables


// Experimental.
type AwsS3TablesTableBucket_IcebergUnreferencedFileRemovalProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table_bucket#settings AwsS3TablesTableBucket#settings}.
	// Experimental.
	Settings *AwsS3TablesTableBucket_SettingsProperty `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table_bucket#status AwsS3TablesTableBucket#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

