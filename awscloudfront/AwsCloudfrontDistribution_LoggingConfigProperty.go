package awscloudfront


// Experimental.
type AwsCloudfrontDistribution_LoggingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#bucket AwsCloudfrontDistribution#bucket}.
	// Experimental.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#include_cookies AwsCloudfrontDistribution#include_cookies}.
	// Experimental.
	IncludeCookies interface{} `field:"optional" json:"includeCookies" yaml:"includeCookies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#prefix AwsCloudfrontDistribution#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

