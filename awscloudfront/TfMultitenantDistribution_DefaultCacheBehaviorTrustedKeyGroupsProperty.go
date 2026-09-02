package awscloudfront


// Experimental.
type TfMultitenantDistribution_DefaultCacheBehaviorTrustedKeyGroupsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#enabled TfMultitenantDistribution#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#items TfMultitenantDistribution#items}.
	// Experimental.
	Items *[]*string `field:"optional" json:"items" yaml:"items"`
}

