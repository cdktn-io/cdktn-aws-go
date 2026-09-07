package kendra


// Experimental.
type AwsDataSource_UrlsProperty struct {
	// seed_url_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#seed_url_configuration AwsDataSource#seed_url_configuration}
	// Experimental.
	SeedUrlConfiguration *AwsDataSource_SeedUrlConfigurationProperty `field:"optional" json:"seedUrlConfiguration" yaml:"seedUrlConfiguration"`
	// site_maps_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#site_maps_configuration AwsDataSource#site_maps_configuration}
	// Experimental.
	SiteMapsConfiguration *AwsDataSource_SiteMapsConfigurationProperty `field:"optional" json:"siteMapsConfiguration" yaml:"siteMapsConfiguration"`
}

