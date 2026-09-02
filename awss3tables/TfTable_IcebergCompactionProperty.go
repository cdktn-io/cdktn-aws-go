package awss3tables


// Experimental.
type TfTable_IcebergCompactionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#settings TfTable#settings}.
	// Experimental.
	Settings *TfTable_MaintenanceConfigurationIcebergCompactionSettingsProperty `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3tables_table#status TfTable#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

