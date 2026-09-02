package awskendra


// Experimental.
type TfDataSource_ConfigurationProperty struct {
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#s3_configuration TfDataSource#s3_configuration}
	// Experimental.
	S3Configuration *TfDataSource_S3ConfigurationProperty `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
	// template_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#template_configuration TfDataSource#template_configuration}
	// Experimental.
	TemplateConfiguration *TfDataSource_TemplateConfigurationProperty `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
	// web_crawler_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#web_crawler_configuration TfDataSource#web_crawler_configuration}
	// Experimental.
	WebCrawlerConfiguration *TfDataSource_WebCrawlerConfigurationProperty `field:"optional" json:"webCrawlerConfiguration" yaml:"webCrawlerConfiguration"`
}

