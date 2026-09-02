package awsopensearch


// Experimental.
type TfPackage_PackageSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_package#s3_bucket_name TfPackage#s3_bucket_name}.
	// Experimental.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_package#s3_key TfPackage#s3_key}.
	// Experimental.
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
}

