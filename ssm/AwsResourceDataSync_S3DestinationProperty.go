package ssm


// Experimental.
type AwsResourceDataSync_S3DestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_resource_data_sync#bucket_name AwsResourceDataSync#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_resource_data_sync#region AwsResourceDataSync#region}.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
	// destination_data_sharing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_resource_data_sync#destination_data_sharing AwsResourceDataSync#destination_data_sharing}
	// Experimental.
	DestinationDataSharing *AwsResourceDataSync_DestinationDataSharingProperty `field:"optional" json:"destinationDataSharing" yaml:"destinationDataSharing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_resource_data_sync#kms_key_arn AwsResourceDataSync#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_resource_data_sync#prefix AwsResourceDataSync#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_resource_data_sync#sync_format AwsResourceDataSync#sync_format}.
	// Experimental.
	SyncFormat *string `field:"optional" json:"syncFormat" yaml:"syncFormat"`
}

