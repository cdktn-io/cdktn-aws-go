package awss3


// Experimental.
type TfBucket_NoncurrentVersionTransitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#storage_class TfBucket#storage_class}.
	// Experimental.
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#days TfBucket#days}.
	// Experimental.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

