package awsdatasync


// Experimental.
type TfTask_S3DestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#bucket_access_role_arn TfTask#bucket_access_role_arn}.
	// Experimental.
	BucketAccessRoleArn *string `field:"required" json:"bucketAccessRoleArn" yaml:"bucketAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#s3_bucket_arn TfTask#s3_bucket_arn}.
	// Experimental.
	S3BucketArn *string `field:"required" json:"s3BucketArn" yaml:"s3BucketArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#subdirectory TfTask#subdirectory}.
	// Experimental.
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
}

