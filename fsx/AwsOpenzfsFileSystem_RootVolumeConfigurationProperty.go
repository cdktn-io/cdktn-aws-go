package fsx


// Experimental.
type AwsOpenzfsFileSystem_RootVolumeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#copy_tags_to_snapshots AwsOpenzfsFileSystem#copy_tags_to_snapshots}.
	// Experimental.
	CopyTagsToSnapshots interface{} `field:"optional" json:"copyTagsToSnapshots" yaml:"copyTagsToSnapshots"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#data_compression_type AwsOpenzfsFileSystem#data_compression_type}.
	// Experimental.
	DataCompressionType *string `field:"optional" json:"dataCompressionType" yaml:"dataCompressionType"`
	// nfs_exports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#nfs_exports AwsOpenzfsFileSystem#nfs_exports}
	// Experimental.
	NfsExports *AwsOpenzfsFileSystem_NfsExportsProperty `field:"optional" json:"nfsExports" yaml:"nfsExports"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#read_only AwsOpenzfsFileSystem#read_only}.
	// Experimental.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#record_size_kib AwsOpenzfsFileSystem#record_size_kib}.
	// Experimental.
	RecordSizeKib *float64 `field:"optional" json:"recordSizeKib" yaml:"recordSizeKib"`
	// user_and_group_quotas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#user_and_group_quotas AwsOpenzfsFileSystem#user_and_group_quotas}
	// Experimental.
	UserAndGroupQuotas interface{} `field:"optional" json:"userAndGroupQuotas" yaml:"userAndGroupQuotas"`
}

