package awsfsx


// Experimental.
type TfDataRepositoryAssociation_S3Property struct {
	// auto_export_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_data_repository_association#auto_export_policy TfDataRepositoryAssociation#auto_export_policy}
	// Experimental.
	AutoExportPolicy *TfDataRepositoryAssociation_AutoExportPolicyProperty `field:"optional" json:"autoExportPolicy" yaml:"autoExportPolicy"`
	// auto_import_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_data_repository_association#auto_import_policy TfDataRepositoryAssociation#auto_import_policy}
	// Experimental.
	AutoImportPolicy *TfDataRepositoryAssociation_AutoImportPolicyProperty `field:"optional" json:"autoImportPolicy" yaml:"autoImportPolicy"`
}

