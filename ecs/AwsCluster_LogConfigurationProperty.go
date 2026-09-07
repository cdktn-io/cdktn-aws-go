package ecs


// Experimental.
type AwsCluster_LogConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#cloud_watch_encryption_enabled AwsCluster#cloud_watch_encryption_enabled}.
	// Experimental.
	CloudWatchEncryptionEnabled interface{} `field:"optional" json:"cloudWatchEncryptionEnabled" yaml:"cloudWatchEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#cloud_watch_log_group_name AwsCluster#cloud_watch_log_group_name}.
	// Experimental.
	CloudWatchLogGroupName *string `field:"optional" json:"cloudWatchLogGroupName" yaml:"cloudWatchLogGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#s3_bucket_encryption_enabled AwsCluster#s3_bucket_encryption_enabled}.
	// Experimental.
	S3BucketEncryptionEnabled interface{} `field:"optional" json:"s3BucketEncryptionEnabled" yaml:"s3BucketEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#s3_bucket_name AwsCluster#s3_bucket_name}.
	// Experimental.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#s3_key_prefix AwsCluster#s3_key_prefix}.
	// Experimental.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

