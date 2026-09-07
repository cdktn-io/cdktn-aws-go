package glue


// Experimental.
type AwsCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#delete_file_threshold AwsCatalogTableOptimizer#delete_file_threshold}.
	// Experimental.
	DeleteFileThreshold *float64 `field:"optional" json:"deleteFileThreshold" yaml:"deleteFileThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#min_input_files AwsCatalogTableOptimizer#min_input_files}.
	// Experimental.
	MinInputFiles *float64 `field:"optional" json:"minInputFiles" yaml:"minInputFiles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#strategy AwsCatalogTableOptimizer#strategy}.
	// Experimental.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

