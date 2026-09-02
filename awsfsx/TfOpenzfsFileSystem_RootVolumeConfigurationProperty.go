package awsfsx


// Experimental.
type TfOpenzfsFileSystem_RootVolumeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#copy_tags_to_snapshots TfOpenzfsFileSystem#copy_tags_to_snapshots}.
	// Experimental.
	CopyTagsToSnapshots interface{} `field:"optional" json:"copyTagsToSnapshots" yaml:"copyTagsToSnapshots"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#data_compression_type TfOpenzfsFileSystem#data_compression_type}.
	// Experimental.
	DataCompressionType *string `field:"optional" json:"dataCompressionType" yaml:"dataCompressionType"`
	// nfs_exports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#nfs_exports TfOpenzfsFileSystem#nfs_exports}
	// Experimental.
	NfsExports *TfOpenzfsFileSystem_NfsExportsProperty `field:"optional" json:"nfsExports" yaml:"nfsExports"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#read_only TfOpenzfsFileSystem#read_only}.
	// Experimental.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#record_size_kib TfOpenzfsFileSystem#record_size_kib}.
	// Experimental.
	RecordSizeKib *float64 `field:"optional" json:"recordSizeKib" yaml:"recordSizeKib"`
	// user_and_group_quotas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#user_and_group_quotas TfOpenzfsFileSystem#user_and_group_quotas}
	// Experimental.
	UserAndGroupQuotas interface{} `field:"optional" json:"userAndGroupQuotas" yaml:"userAndGroupQuotas"`
}

