package awss3tables


// Experimental.
type AwsS3TablesTable_MaintenanceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#iceberg_compaction AwsS3TablesTable#iceberg_compaction}.
	// Experimental.
	IcebergCompaction *AwsS3TablesTable_IcebergCompactionProperty `field:"optional" json:"icebergCompaction" yaml:"icebergCompaction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#iceberg_snapshot_management AwsS3TablesTable#iceberg_snapshot_management}.
	// Experimental.
	IcebergSnapshotManagement *AwsS3TablesTable_IcebergSnapshotManagementProperty `field:"optional" json:"icebergSnapshotManagement" yaml:"icebergSnapshotManagement"`
}

