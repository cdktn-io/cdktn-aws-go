package awsdms


// Experimental.
type TfEndpoint_RedshiftSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#bucket_folder TfEndpoint#bucket_folder}.
	// Experimental.
	BucketFolder *string `field:"optional" json:"bucketFolder" yaml:"bucketFolder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#bucket_name TfEndpoint#bucket_name}.
	// Experimental.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#encryption_mode TfEndpoint#encryption_mode}.
	// Experimental.
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#server_side_encryption_kms_key_id TfEndpoint#server_side_encryption_kms_key_id}.
	// Experimental.
	ServerSideEncryptionKmsKeyId *string `field:"optional" json:"serverSideEncryptionKmsKeyId" yaml:"serverSideEncryptionKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#service_access_role_arn TfEndpoint#service_access_role_arn}.
	// Experimental.
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
}

