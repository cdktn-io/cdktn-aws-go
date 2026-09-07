package efs


// Experimental.
type AwsReplicationConfiguration_DestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_replication_configuration#availability_zone_name AwsReplicationConfiguration#availability_zone_name}.
	// Experimental.
	AvailabilityZoneName *string `field:"optional" json:"availabilityZoneName" yaml:"availabilityZoneName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_replication_configuration#file_system_id AwsReplicationConfiguration#file_system_id}.
	// Experimental.
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_replication_configuration#kms_key_id AwsReplicationConfiguration#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_replication_configuration#region AwsReplicationConfiguration#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

