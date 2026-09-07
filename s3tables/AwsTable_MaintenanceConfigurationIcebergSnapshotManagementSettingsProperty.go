package s3tables


// Experimental.
type AwsTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#max_snapshot_age_hours AwsTable#max_snapshot_age_hours}.
	// Experimental.
	MaxSnapshotAgeHours *float64 `field:"optional" json:"maxSnapshotAgeHours" yaml:"maxSnapshotAgeHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#min_snapshots_to_keep AwsTable#min_snapshots_to_keep}.
	// Experimental.
	MinSnapshotsToKeep *float64 `field:"optional" json:"minSnapshotsToKeep" yaml:"minSnapshotsToKeep"`
}

