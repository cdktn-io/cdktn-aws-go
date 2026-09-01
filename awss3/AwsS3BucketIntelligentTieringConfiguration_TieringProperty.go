package awss3


// Experimental.
type AwsS3BucketIntelligentTieringConfiguration_TieringProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_intelligent_tiering_configuration#access_tier AwsS3BucketIntelligentTieringConfiguration#access_tier}.
	// Experimental.
	AccessTier *string `field:"required" json:"accessTier" yaml:"accessTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_intelligent_tiering_configuration#days AwsS3BucketIntelligentTieringConfiguration#days}.
	// Experimental.
	Days *float64 `field:"required" json:"days" yaml:"days"`
}

