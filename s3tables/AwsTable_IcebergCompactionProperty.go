package s3tables


// Experimental.
type AwsTable_IcebergCompactionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#settings AwsTable#settings}.
	// Experimental.
	Settings *AwsTable_MaintenanceConfigurationIcebergCompactionSettingsProperty `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#status AwsTable#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

