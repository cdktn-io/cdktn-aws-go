package awss3


// Experimental.
type TfBucketLifecycleConfiguration_FilterProperty struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#and TfBucketLifecycleConfiguration#and}
	// Experimental.
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#object_size_greater_than TfBucketLifecycleConfiguration#object_size_greater_than}.
	// Experimental.
	ObjectSizeGreaterThan *float64 `field:"optional" json:"objectSizeGreaterThan" yaml:"objectSizeGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#object_size_less_than TfBucketLifecycleConfiguration#object_size_less_than}.
	// Experimental.
	ObjectSizeLessThan *float64 `field:"optional" json:"objectSizeLessThan" yaml:"objectSizeLessThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#prefix TfBucketLifecycleConfiguration#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#tag TfBucketLifecycleConfiguration#tag}
	// Experimental.
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

