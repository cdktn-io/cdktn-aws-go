package awss3tables


// Experimental.
type TfTable_MaintenanceConfigurationIcebergSnapshotManagementSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#max_snapshot_age_hours TfTable#max_snapshot_age_hours}.
	// Experimental.
	MaxSnapshotAgeHours *float64 `field:"optional" json:"maxSnapshotAgeHours" yaml:"maxSnapshotAgeHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#min_snapshots_to_keep TfTable#min_snapshots_to_keep}.
	// Experimental.
	MinSnapshotsToKeep *float64 `field:"optional" json:"minSnapshotsToKeep" yaml:"minSnapshotsToKeep"`
}

