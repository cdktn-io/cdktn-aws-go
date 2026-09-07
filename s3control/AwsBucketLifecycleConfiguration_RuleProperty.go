package s3control


// Experimental.
type AwsBucketLifecycleConfiguration_RuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_bucket_lifecycle_configuration#id AwsBucketLifecycleConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// abort_incomplete_multipart_upload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_bucket_lifecycle_configuration#abort_incomplete_multipart_upload AwsBucketLifecycleConfiguration#abort_incomplete_multipart_upload}
	// Experimental.
	AbortIncompleteMultipartUpload *AwsBucketLifecycleConfiguration_AbortIncompleteMultipartUploadProperty `field:"optional" json:"abortIncompleteMultipartUpload" yaml:"abortIncompleteMultipartUpload"`
	// expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_bucket_lifecycle_configuration#expiration AwsBucketLifecycleConfiguration#expiration}
	// Experimental.
	Expiration *AwsBucketLifecycleConfiguration_ExpirationProperty `field:"optional" json:"expiration" yaml:"expiration"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_bucket_lifecycle_configuration#filter AwsBucketLifecycleConfiguration#filter}
	// Experimental.
	Filter *AwsBucketLifecycleConfiguration_FilterProperty `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_bucket_lifecycle_configuration#status AwsBucketLifecycleConfiguration#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

