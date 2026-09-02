package awss3


// Experimental.
type TfBucketLifecycleConfiguration_TagProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#key TfBucketLifecycleConfiguration#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#value TfBucketLifecycleConfiguration#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

