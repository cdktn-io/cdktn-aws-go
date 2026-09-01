package awsfsx


// Experimental.
type AwsFsxFileCache_LustreConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_file_cache#deployment_type AwsFsxFileCache#deployment_type}.
	// Experimental.
	DeploymentType *string `field:"required" json:"deploymentType" yaml:"deploymentType"`
	// metadata_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_file_cache#metadata_configuration AwsFsxFileCache#metadata_configuration}
	// Experimental.
	MetadataConfiguration interface{} `field:"required" json:"metadataConfiguration" yaml:"metadataConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_file_cache#per_unit_storage_throughput AwsFsxFileCache#per_unit_storage_throughput}.
	// Experimental.
	PerUnitStorageThroughput *float64 `field:"required" json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_file_cache#weekly_maintenance_start_time AwsFsxFileCache#weekly_maintenance_start_time}.
	// Experimental.
	WeeklyMaintenanceStartTime *string `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

