package awss3


// Experimental.
type TfBucketLifecycleConfiguration_TransitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#storage_class TfBucketLifecycleConfiguration#storage_class}.
	// Experimental.
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#date TfBucketLifecycleConfiguration#date}.
	// Experimental.
	Date *string `field:"optional" json:"date" yaml:"date"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#days TfBucketLifecycleConfiguration#days}.
	// Experimental.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

