package awsappsync


// Experimental.
type TfResolver_CachingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#caching_keys TfResolver#caching_keys}.
	// Experimental.
	CachingKeys *[]*string `field:"optional" json:"cachingKeys" yaml:"cachingKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#ttl TfResolver#ttl}.
	// Experimental.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

