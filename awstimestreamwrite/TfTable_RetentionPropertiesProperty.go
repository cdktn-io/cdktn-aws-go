package awstimestreamwrite


// Experimental.
type TfTable_RetentionPropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#magnetic_store_retention_period_in_days TfTable#magnetic_store_retention_period_in_days}.
	// Experimental.
	MagneticStoreRetentionPeriodInDays *float64 `field:"required" json:"magneticStoreRetentionPeriodInDays" yaml:"magneticStoreRetentionPeriodInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamwrite_table#memory_store_retention_period_in_hours TfTable#memory_store_retention_period_in_hours}.
	// Experimental.
	MemoryStoreRetentionPeriodInHours *float64 `field:"required" json:"memoryStoreRetentionPeriodInHours" yaml:"memoryStoreRetentionPeriodInHours"`
}

