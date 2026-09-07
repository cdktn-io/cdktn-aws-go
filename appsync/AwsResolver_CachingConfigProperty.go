package appsync


// Experimental.
type AwsResolver_CachingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#caching_keys AwsResolver#caching_keys}.
	// Experimental.
	CachingKeys *[]*string `field:"optional" json:"cachingKeys" yaml:"cachingKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_resolver#ttl AwsResolver#ttl}.
	// Experimental.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

