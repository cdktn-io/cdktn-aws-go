package cloudfront


// Experimental.
type AwsMultitenantDistribution_CacheBehaviorTrustedKeyGroupsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#enabled AwsMultitenantDistribution#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#items AwsMultitenantDistribution#items}.
	// Experimental.
	Items *[]*string `field:"optional" json:"items" yaml:"items"`
}

