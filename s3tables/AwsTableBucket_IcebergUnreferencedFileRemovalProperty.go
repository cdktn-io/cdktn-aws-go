package s3tables


// Experimental.
type AwsTableBucket_IcebergUnreferencedFileRemovalProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table_bucket#settings AwsTableBucket#settings}.
	// Experimental.
	Settings *AwsTableBucket_SettingsProperty `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table_bucket#status AwsTableBucket#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

