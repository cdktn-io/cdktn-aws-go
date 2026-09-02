package awss3


// Experimental.
type TfBucketMetadataConfiguration_MetadataConfigurationProperty struct {
	// inventory_table_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_metadata_configuration#inventory_table_configuration TfBucketMetadataConfiguration#inventory_table_configuration}
	// Experimental.
	InventoryTableConfiguration interface{} `field:"optional" json:"inventoryTableConfiguration" yaml:"inventoryTableConfiguration"`
	// journal_table_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_metadata_configuration#journal_table_configuration TfBucketMetadataConfiguration#journal_table_configuration}
	// Experimental.
	JournalTableConfiguration interface{} `field:"optional" json:"journalTableConfiguration" yaml:"journalTableConfiguration"`
}

