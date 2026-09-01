package awsstoragegateway


// Experimental.
type AwsStoragegatewayNfsFileShare_NfsFileShareDefaultsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#directory_mode AwsStoragegatewayNfsFileShare#directory_mode}.
	// Experimental.
	DirectoryMode *string `field:"optional" json:"directoryMode" yaml:"directoryMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#file_mode AwsStoragegatewayNfsFileShare#file_mode}.
	// Experimental.
	FileMode *string `field:"optional" json:"fileMode" yaml:"fileMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#group_id AwsStoragegatewayNfsFileShare#group_id}.
	// Experimental.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#owner_id AwsStoragegatewayNfsFileShare#owner_id}.
	// Experimental.
	OwnerId *string `field:"optional" json:"ownerId" yaml:"ownerId"`
}

