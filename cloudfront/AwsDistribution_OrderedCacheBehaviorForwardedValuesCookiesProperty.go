package cloudfront


// Experimental.
type AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#forward AwsDistribution#forward}.
	// Experimental.
	Forward *string `field:"required" json:"forward" yaml:"forward"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#whitelisted_names AwsDistribution#whitelisted_names}.
	// Experimental.
	WhitelistedNames *[]*string `field:"optional" json:"whitelistedNames" yaml:"whitelistedNames"`
}

