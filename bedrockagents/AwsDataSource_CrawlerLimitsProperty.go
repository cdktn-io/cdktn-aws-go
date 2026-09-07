package bedrockagents


// Experimental.
type AwsDataSource_CrawlerLimitsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#max_pages AwsDataSource#max_pages}.
	// Experimental.
	MaxPages *float64 `field:"optional" json:"maxPages" yaml:"maxPages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#rate_limit AwsDataSource#rate_limit}.
	// Experimental.
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

