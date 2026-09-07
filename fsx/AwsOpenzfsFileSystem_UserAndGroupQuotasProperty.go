package fsx


// Experimental.
type AwsOpenzfsFileSystem_UserAndGroupQuotasProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#id AwsOpenzfsFileSystem#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *float64 `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#storage_capacity_quota_gib AwsOpenzfsFileSystem#storage_capacity_quota_gib}.
	// Experimental.
	StorageCapacityQuotaGib *float64 `field:"required" json:"storageCapacityQuotaGib" yaml:"storageCapacityQuotaGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#type AwsOpenzfsFileSystem#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

