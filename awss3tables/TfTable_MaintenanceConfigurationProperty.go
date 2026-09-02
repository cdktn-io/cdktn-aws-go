package awss3tables


// Experimental.
type TfTable_MaintenanceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#iceberg_compaction TfTable#iceberg_compaction}.
	// Experimental.
	IcebergCompaction *TfTable_IcebergCompactionProperty `field:"optional" json:"icebergCompaction" yaml:"icebergCompaction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#iceberg_snapshot_management TfTable#iceberg_snapshot_management}.
	// Experimental.
	IcebergSnapshotManagement *TfTable_IcebergSnapshotManagementProperty `field:"optional" json:"icebergSnapshotManagement" yaml:"icebergSnapshotManagement"`
}

