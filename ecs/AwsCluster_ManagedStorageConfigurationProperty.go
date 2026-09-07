package ecs


// Experimental.
type AwsCluster_ManagedStorageConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#fargate_ephemeral_storage_kms_key_id AwsCluster#fargate_ephemeral_storage_kms_key_id}.
	// Experimental.
	FargateEphemeralStorageKmsKeyId *string `field:"optional" json:"fargateEphemeralStorageKmsKeyId" yaml:"fargateEphemeralStorageKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#kms_key_id AwsCluster#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

