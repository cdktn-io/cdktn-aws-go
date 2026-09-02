package awssecuritylake


// Experimental.
type TfCustomLogSource_ConfigurationProperty struct {
	// crawler_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_custom_log_source#crawler_configuration TfCustomLogSource#crawler_configuration}
	// Experimental.
	CrawlerConfiguration interface{} `field:"optional" json:"crawlerConfiguration" yaml:"crawlerConfiguration"`
	// provider_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_custom_log_source#provider_identity TfCustomLogSource#provider_identity}
	// Experimental.
	ProviderIdentity interface{} `field:"optional" json:"providerIdentity" yaml:"providerIdentity"`
}

