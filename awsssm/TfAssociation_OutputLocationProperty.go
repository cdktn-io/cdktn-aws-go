package awsssm


// Experimental.
type TfAssociation_OutputLocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_association#s3_bucket_name TfAssociation#s3_bucket_name}.
	// Experimental.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_association#s3_key_prefix TfAssociation#s3_key_prefix}.
	// Experimental.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssm_association#s3_region TfAssociation#s3_region}.
	// Experimental.
	S3Region *string `field:"optional" json:"s3Region" yaml:"s3Region"`
}

