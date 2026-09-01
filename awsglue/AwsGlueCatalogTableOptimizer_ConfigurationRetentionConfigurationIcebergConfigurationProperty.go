package awsglue


// Experimental.
type AwsGlueCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#clean_expired_files AwsGlueCatalogTableOptimizer#clean_expired_files}.
	// Experimental.
	CleanExpiredFiles interface{} `field:"optional" json:"cleanExpiredFiles" yaml:"cleanExpiredFiles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#number_of_snapshots_to_retain AwsGlueCatalogTableOptimizer#number_of_snapshots_to_retain}.
	// Experimental.
	NumberOfSnapshotsToRetain *float64 `field:"optional" json:"numberOfSnapshotsToRetain" yaml:"numberOfSnapshotsToRetain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#run_rate_in_hours AwsGlueCatalogTableOptimizer#run_rate_in_hours}.
	// Experimental.
	RunRateInHours *float64 `field:"optional" json:"runRateInHours" yaml:"runRateInHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#snapshot_retention_period_in_days AwsGlueCatalogTableOptimizer#snapshot_retention_period_in_days}.
	// Experimental.
	SnapshotRetentionPeriodInDays *float64 `field:"optional" json:"snapshotRetentionPeriodInDays" yaml:"snapshotRetentionPeriodInDays"`
}

