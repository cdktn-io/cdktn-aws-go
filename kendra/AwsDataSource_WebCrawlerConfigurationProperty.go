package kendra


// Experimental.
type AwsDataSource_WebCrawlerConfigurationProperty struct {
	// urls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#urls AwsDataSource#urls}
	// Experimental.
	Urls *AwsDataSource_UrlsProperty `field:"required" json:"urls" yaml:"urls"`
	// authentication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#authentication_configuration AwsDataSource#authentication_configuration}
	// Experimental.
	AuthenticationConfiguration *AwsDataSource_AuthenticationConfigurationProperty `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#crawl_depth AwsDataSource#crawl_depth}.
	// Experimental.
	CrawlDepth *float64 `field:"optional" json:"crawlDepth" yaml:"crawlDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#max_content_size_per_page_in_mega_bytes AwsDataSource#max_content_size_per_page_in_mega_bytes}.
	// Experimental.
	MaxContentSizePerPageInMegaBytes *float64 `field:"optional" json:"maxContentSizePerPageInMegaBytes" yaml:"maxContentSizePerPageInMegaBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#max_links_per_page AwsDataSource#max_links_per_page}.
	// Experimental.
	MaxLinksPerPage *float64 `field:"optional" json:"maxLinksPerPage" yaml:"maxLinksPerPage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#max_urls_per_minute_crawl_rate AwsDataSource#max_urls_per_minute_crawl_rate}.
	// Experimental.
	MaxUrlsPerMinuteCrawlRate *float64 `field:"optional" json:"maxUrlsPerMinuteCrawlRate" yaml:"maxUrlsPerMinuteCrawlRate"`
	// proxy_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#proxy_configuration AwsDataSource#proxy_configuration}
	// Experimental.
	ProxyConfiguration *AwsDataSource_ProxyConfigurationProperty `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#url_exclusion_patterns AwsDataSource#url_exclusion_patterns}.
	// Experimental.
	UrlExclusionPatterns *[]*string `field:"optional" json:"urlExclusionPatterns" yaml:"urlExclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#url_inclusion_patterns AwsDataSource#url_inclusion_patterns}.
	// Experimental.
	UrlInclusionPatterns *[]*string `field:"optional" json:"urlInclusionPatterns" yaml:"urlInclusionPatterns"`
}

