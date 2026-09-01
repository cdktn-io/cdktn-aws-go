package awskendra


// Experimental.
type AwsKendraDataSource_SeedUrlConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#seed_urls AwsKendraDataSource#seed_urls}.
	// Experimental.
	SeedUrls *[]*string `field:"required" json:"seedUrls" yaml:"seedUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#web_crawler_mode AwsKendraDataSource#web_crawler_mode}.
	// Experimental.
	WebCrawlerMode *string `field:"optional" json:"webCrawlerMode" yaml:"webCrawlerMode"`
}

