package awss3


// Experimental.
type TfBucket_LoggingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#target_bucket TfBucket#target_bucket}.
	// Experimental.
	TargetBucket *string `field:"required" json:"targetBucket" yaml:"targetBucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#target_prefix TfBucket#target_prefix}.
	// Experimental.
	TargetPrefix *string `field:"optional" json:"targetPrefix" yaml:"targetPrefix"`
}

