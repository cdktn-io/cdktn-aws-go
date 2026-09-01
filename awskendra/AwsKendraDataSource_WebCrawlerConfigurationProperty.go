package awskendra


// Experimental.
type AwsKendraDataSource_WebCrawlerConfigurationProperty struct {
	// urls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#urls AwsKendraDataSource#urls}
	// Experimental.
	Urls *AwsKendraDataSource_UrlsProperty `field:"required" json:"urls" yaml:"urls"`
	// authentication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#authentication_configuration AwsKendraDataSource#authentication_configuration}
	// Experimental.
	AuthenticationConfiguration *AwsKendraDataSource_AuthenticationConfigurationProperty `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#crawl_depth AwsKendraDataSource#crawl_depth}.
	// Experimental.
	CrawlDepth *float64 `field:"optional" json:"crawlDepth" yaml:"crawlDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#max_content_size_per_page_in_mega_bytes AwsKendraDataSource#max_content_size_per_page_in_mega_bytes}.
	// Experimental.
	MaxContentSizePerPageInMegaBytes *float64 `field:"optional" json:"maxContentSizePerPageInMegaBytes" yaml:"maxContentSizePerPageInMegaBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#max_links_per_page AwsKendraDataSource#max_links_per_page}.
	// Experimental.
	MaxLinksPerPage *float64 `field:"optional" json:"maxLinksPerPage" yaml:"maxLinksPerPage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#max_urls_per_minute_crawl_rate AwsKendraDataSource#max_urls_per_minute_crawl_rate}.
	// Experimental.
	MaxUrlsPerMinuteCrawlRate *float64 `field:"optional" json:"maxUrlsPerMinuteCrawlRate" yaml:"maxUrlsPerMinuteCrawlRate"`
	// proxy_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#proxy_configuration AwsKendraDataSource#proxy_configuration}
	// Experimental.
	ProxyConfiguration *AwsKendraDataSource_ProxyConfigurationProperty `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#url_exclusion_patterns AwsKendraDataSource#url_exclusion_patterns}.
	// Experimental.
	UrlExclusionPatterns *[]*string `field:"optional" json:"urlExclusionPatterns" yaml:"urlExclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#url_inclusion_patterns AwsKendraDataSource#url_inclusion_patterns}.
	// Experimental.
	UrlInclusionPatterns *[]*string `field:"optional" json:"urlInclusionPatterns" yaml:"urlInclusionPatterns"`
}

