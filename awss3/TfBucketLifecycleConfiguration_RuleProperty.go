package awss3


// Experimental.
type TfBucketLifecycleConfiguration_RuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#id TfBucketLifecycleConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#status TfBucketLifecycleConfiguration#status}.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
	// abort_incomplete_multipart_upload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#abort_incomplete_multipart_upload TfBucketLifecycleConfiguration#abort_incomplete_multipart_upload}
	// Experimental.
	AbortIncompleteMultipartUpload interface{} `field:"optional" json:"abortIncompleteMultipartUpload" yaml:"abortIncompleteMultipartUpload"`
	// expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#expiration TfBucketLifecycleConfiguration#expiration}
	// Experimental.
	Expiration interface{} `field:"optional" json:"expiration" yaml:"expiration"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#filter TfBucketLifecycleConfiguration#filter}
	// Experimental.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// noncurrent_version_expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#noncurrent_version_expiration TfBucketLifecycleConfiguration#noncurrent_version_expiration}
	// Experimental.
	NoncurrentVersionExpiration interface{} `field:"optional" json:"noncurrentVersionExpiration" yaml:"noncurrentVersionExpiration"`
	// noncurrent_version_transition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#noncurrent_version_transition TfBucketLifecycleConfiguration#noncurrent_version_transition}
	// Experimental.
	NoncurrentVersionTransition interface{} `field:"optional" json:"noncurrentVersionTransition" yaml:"noncurrentVersionTransition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#prefix TfBucketLifecycleConfiguration#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// transition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#transition TfBucketLifecycleConfiguration#transition}
	// Experimental.
	Transition interface{} `field:"optional" json:"transition" yaml:"transition"`
}

