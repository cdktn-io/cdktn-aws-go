package awskendra


// Experimental.
type TfDataSource_UrlsProperty struct {
	// seed_url_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#seed_url_configuration TfDataSource#seed_url_configuration}
	// Experimental.
	SeedUrlConfiguration *TfDataSource_SeedUrlConfigurationProperty `field:"optional" json:"seedUrlConfiguration" yaml:"seedUrlConfiguration"`
	// site_maps_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#site_maps_configuration TfDataSource#site_maps_configuration}
	// Experimental.
	SiteMapsConfiguration *TfDataSource_SiteMapsConfigurationProperty `field:"optional" json:"siteMapsConfiguration" yaml:"siteMapsConfiguration"`
}

