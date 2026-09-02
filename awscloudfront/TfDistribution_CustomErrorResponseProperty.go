package awscloudfront


// Experimental.
type TfDistribution_CustomErrorResponseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#error_code TfDistribution#error_code}.
	// Experimental.
	ErrorCode *float64 `field:"required" json:"errorCode" yaml:"errorCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#error_caching_min_ttl TfDistribution#error_caching_min_ttl}.
	// Experimental.
	ErrorCachingMinTtl *float64 `field:"optional" json:"errorCachingMinTtl" yaml:"errorCachingMinTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#response_code TfDistribution#response_code}.
	// Experimental.
	ResponseCode *float64 `field:"optional" json:"responseCode" yaml:"responseCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#response_page_path TfDistribution#response_page_path}.
	// Experimental.
	ResponsePagePath *string `field:"optional" json:"responsePagePath" yaml:"responsePagePath"`
}

