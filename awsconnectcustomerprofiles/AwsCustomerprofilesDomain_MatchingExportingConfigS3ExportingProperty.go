package awsconnectcustomerprofiles


// Experimental.
type AwsCustomerprofilesDomain_MatchingExportingConfigS3ExportingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#s3_bucket_name AwsCustomerprofilesDomain#s3_bucket_name}.
	// Experimental.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#s3_key_name AwsCustomerprofilesDomain#s3_key_name}.
	// Experimental.
	S3KeyName *string `field:"optional" json:"s3KeyName" yaml:"s3KeyName"`
}

