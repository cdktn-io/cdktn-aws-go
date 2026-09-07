package s3tables


// Experimental.
type AwsTable_MaintenanceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#iceberg_compaction AwsTable#iceberg_compaction}.
	// Experimental.
	IcebergCompaction *AwsTable_IcebergCompactionProperty `field:"optional" json:"icebergCompaction" yaml:"icebergCompaction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#iceberg_snapshot_management AwsTable#iceberg_snapshot_management}.
	// Experimental.
	IcebergSnapshotManagement *AwsTable_IcebergSnapshotManagementProperty `field:"optional" json:"icebergSnapshotManagement" yaml:"icebergSnapshotManagement"`
}

