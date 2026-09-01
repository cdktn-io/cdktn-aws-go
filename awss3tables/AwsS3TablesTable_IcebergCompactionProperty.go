package awss3tables


// Experimental.
type AwsS3TablesTable_IcebergCompactionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#settings AwsS3TablesTable#settings}.
	// Experimental.
	Settings *AwsS3TablesTable_MaintenanceConfigurationIcebergCompactionSettingsProperty `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#status AwsS3TablesTable#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

