package awss3tables


// Experimental.
type TfTableBucket_IcebergUnreferencedFileRemovalProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table_bucket#settings TfTableBucket#settings}.
	// Experimental.
	Settings *TfTableBucket_SettingsProperty `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table_bucket#status TfTableBucket#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

