package s3


// Experimental.
type AwsBucketIntelligentTieringConfiguration_TieringProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_intelligent_tiering_configuration#access_tier AwsBucketIntelligentTieringConfiguration#access_tier}.
	// Experimental.
	AccessTier *string `field:"required" json:"accessTier" yaml:"accessTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_intelligent_tiering_configuration#days AwsBucketIntelligentTieringConfiguration#days}.
	// Experimental.
	Days *float64 `field:"required" json:"days" yaml:"days"`
}

