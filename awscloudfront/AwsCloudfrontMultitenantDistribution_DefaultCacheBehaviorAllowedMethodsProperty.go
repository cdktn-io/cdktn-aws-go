package awscloudfront


// Experimental.
type AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorAllowedMethodsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#cached_methods AwsCloudfrontMultitenantDistribution#cached_methods}.
	// Experimental.
	CachedMethods *[]*string `field:"required" json:"cachedMethods" yaml:"cachedMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#items AwsCloudfrontMultitenantDistribution#items}.
	// Experimental.
	Items *[]*string `field:"required" json:"items" yaml:"items"`
}

