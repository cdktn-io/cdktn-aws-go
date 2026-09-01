package awskendra


// Experimental.
type AwsKendraDataSource_UrlsProperty struct {
	// seed_url_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#seed_url_configuration AwsKendraDataSource#seed_url_configuration}
	// Experimental.
	SeedUrlConfiguration *AwsKendraDataSource_SeedUrlConfigurationProperty `field:"optional" json:"seedUrlConfiguration" yaml:"seedUrlConfiguration"`
	// site_maps_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#site_maps_configuration AwsKendraDataSource#site_maps_configuration}
	// Experimental.
	SiteMapsConfiguration *AwsKendraDataSource_SiteMapsConfigurationProperty `field:"optional" json:"siteMapsConfiguration" yaml:"siteMapsConfiguration"`
}

