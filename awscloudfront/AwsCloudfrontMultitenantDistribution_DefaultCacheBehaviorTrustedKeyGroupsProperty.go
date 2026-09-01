package awscloudfront


// Experimental.
type AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorTrustedKeyGroupsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#enabled AwsCloudfrontMultitenantDistribution#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#items AwsCloudfrontMultitenantDistribution#items}.
	// Experimental.
	Items *[]*string `field:"optional" json:"items" yaml:"items"`
}

