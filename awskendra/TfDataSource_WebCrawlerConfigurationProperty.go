package awskendra


// Experimental.
type TfDataSource_WebCrawlerConfigurationProperty struct {
	// urls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#urls TfDataSource#urls}
	// Experimental.
	Urls *TfDataSource_UrlsProperty `field:"required" json:"urls" yaml:"urls"`
	// authentication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#authentication_configuration TfDataSource#authentication_configuration}
	// Experimental.
	AuthenticationConfiguration *TfDataSource_AuthenticationConfigurationProperty `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#crawl_depth TfDataSource#crawl_depth}.
	// Experimental.
	CrawlDepth *float64 `field:"optional" json:"crawlDepth" yaml:"crawlDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#max_content_size_per_page_in_mega_bytes TfDataSource#max_content_size_per_page_in_mega_bytes}.
	// Experimental.
	MaxContentSizePerPageInMegaBytes *float64 `field:"optional" json:"maxContentSizePerPageInMegaBytes" yaml:"maxContentSizePerPageInMegaBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#max_links_per_page TfDataSource#max_links_per_page}.
	// Experimental.
	MaxLinksPerPage *float64 `field:"optional" json:"maxLinksPerPage" yaml:"maxLinksPerPage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#max_urls_per_minute_crawl_rate TfDataSource#max_urls_per_minute_crawl_rate}.
	// Experimental.
	MaxUrlsPerMinuteCrawlRate *float64 `field:"optional" json:"maxUrlsPerMinuteCrawlRate" yaml:"maxUrlsPerMinuteCrawlRate"`
	// proxy_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#proxy_configuration TfDataSource#proxy_configuration}
	// Experimental.
	ProxyConfiguration *TfDataSource_ProxyConfigurationProperty `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#url_exclusion_patterns TfDataSource#url_exclusion_patterns}.
	// Experimental.
	UrlExclusionPatterns *[]*string `field:"optional" json:"urlExclusionPatterns" yaml:"urlExclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#url_inclusion_patterns TfDataSource#url_inclusion_patterns}.
	// Experimental.
	UrlInclusionPatterns *[]*string `field:"optional" json:"urlInclusionPatterns" yaml:"urlInclusionPatterns"`
}

