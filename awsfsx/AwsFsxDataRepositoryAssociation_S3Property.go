package awsfsx


// Experimental.
type AwsFsxDataRepositoryAssociation_S3Property struct {
	// auto_export_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_data_repository_association#auto_export_policy AwsFsxDataRepositoryAssociation#auto_export_policy}
	// Experimental.
	AutoExportPolicy *AwsFsxDataRepositoryAssociation_AutoExportPolicyProperty `field:"optional" json:"autoExportPolicy" yaml:"autoExportPolicy"`
	// auto_import_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_data_repository_association#auto_import_policy AwsFsxDataRepositoryAssociation#auto_import_policy}
	// Experimental.
	AutoImportPolicy *AwsFsxDataRepositoryAssociation_AutoImportPolicyProperty `field:"optional" json:"autoImportPolicy" yaml:"autoImportPolicy"`
}

