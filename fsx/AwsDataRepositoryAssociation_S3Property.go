package fsx


// Experimental.
type AwsDataRepositoryAssociation_S3Property struct {
	// auto_export_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_data_repository_association#auto_export_policy AwsDataRepositoryAssociation#auto_export_policy}
	// Experimental.
	AutoExportPolicy *AwsDataRepositoryAssociation_AutoExportPolicyProperty `field:"optional" json:"autoExportPolicy" yaml:"autoExportPolicy"`
	// auto_import_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_data_repository_association#auto_import_policy AwsDataRepositoryAssociation#auto_import_policy}
	// Experimental.
	AutoImportPolicy *AwsDataRepositoryAssociation_AutoImportPolicyProperty `field:"optional" json:"autoImportPolicy" yaml:"autoImportPolicy"`
}

