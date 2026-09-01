package awscloudfront


// Experimental.
type AwsCloudfrontMultitenantDistribution_CustomErrorResponseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#error_code AwsCloudfrontMultitenantDistribution#error_code}.
	// Experimental.
	ErrorCode *float64 `field:"required" json:"errorCode" yaml:"errorCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#error_caching_min_ttl AwsCloudfrontMultitenantDistribution#error_caching_min_ttl}.
	// Experimental.
	ErrorCachingMinTtl *float64 `field:"optional" json:"errorCachingMinTtl" yaml:"errorCachingMinTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#response_code AwsCloudfrontMultitenantDistribution#response_code}.
	// Experimental.
	ResponseCode *string `field:"optional" json:"responseCode" yaml:"responseCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#response_page_path AwsCloudfrontMultitenantDistribution#response_page_path}.
	// Experimental.
	ResponsePagePath *string `field:"optional" json:"responsePagePath" yaml:"responsePagePath"`
}

