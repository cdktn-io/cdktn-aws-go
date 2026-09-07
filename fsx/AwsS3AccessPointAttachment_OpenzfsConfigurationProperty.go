package fsx


// Experimental.
type AwsS3AccessPointAttachment_OpenzfsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_s3_access_point_attachment#volume_id AwsS3AccessPointAttachment#volume_id}.
	// Experimental.
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// file_system_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_s3_access_point_attachment#file_system_identity AwsS3AccessPointAttachment#file_system_identity}
	// Experimental.
	FileSystemIdentity interface{} `field:"optional" json:"fileSystemIdentity" yaml:"fileSystemIdentity"`
}

