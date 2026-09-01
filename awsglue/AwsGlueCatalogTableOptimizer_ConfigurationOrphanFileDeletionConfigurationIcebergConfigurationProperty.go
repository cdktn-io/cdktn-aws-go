package awsglue


// Experimental.
type AwsGlueCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#location AwsGlueCatalogTableOptimizer#location}.
	// Experimental.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#orphan_file_retention_period_in_days AwsGlueCatalogTableOptimizer#orphan_file_retention_period_in_days}.
	// Experimental.
	OrphanFileRetentionPeriodInDays *float64 `field:"optional" json:"orphanFileRetentionPeriodInDays" yaml:"orphanFileRetentionPeriodInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#run_rate_in_hours AwsGlueCatalogTableOptimizer#run_rate_in_hours}.
	// Experimental.
	RunRateInHours *float64 `field:"optional" json:"runRateInHours" yaml:"runRateInHours"`
}

