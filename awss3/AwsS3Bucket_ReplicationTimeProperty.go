package awss3


// Experimental.
type AwsS3Bucket_ReplicationTimeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#minutes AwsS3Bucket#minutes}.
	// Experimental.
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#status AwsS3Bucket#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

