package awscloudfront


// Experimental.
type AwsCloudfrontDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#forward AwsCloudfrontDistribution#forward}.
	// Experimental.
	Forward *string `field:"required" json:"forward" yaml:"forward"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#whitelisted_names AwsCloudfrontDistribution#whitelisted_names}.
	// Experimental.
	WhitelistedNames *[]*string `field:"optional" json:"whitelistedNames" yaml:"whitelistedNames"`
}

