package awsglue


// Experimental.
type TfCatalogTableOptimizer_ConfigurationCompactionConfigurationIcebergConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#delete_file_threshold TfCatalogTableOptimizer#delete_file_threshold}.
	// Experimental.
	DeleteFileThreshold *float64 `field:"optional" json:"deleteFileThreshold" yaml:"deleteFileThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#min_input_files TfCatalogTableOptimizer#min_input_files}.
	// Experimental.
	MinInputFiles *float64 `field:"optional" json:"minInputFiles" yaml:"minInputFiles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#strategy TfCatalogTableOptimizer#strategy}.
	// Experimental.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

