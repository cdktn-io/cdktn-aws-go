package awskendra


// Experimental.
type AwsKendraDataSource_ConfigurationProperty struct {
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#s3_configuration AwsKendraDataSource#s3_configuration}
	// Experimental.
	S3Configuration *AwsKendraDataSource_S3ConfigurationProperty `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
	// template_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#template_configuration AwsKendraDataSource#template_configuration}
	// Experimental.
	TemplateConfiguration *AwsKendraDataSource_TemplateConfigurationProperty `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
	// web_crawler_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#web_crawler_configuration AwsKendraDataSource#web_crawler_configuration}
	// Experimental.
	WebCrawlerConfiguration *AwsKendraDataSource_WebCrawlerConfigurationProperty `field:"optional" json:"webCrawlerConfiguration" yaml:"webCrawlerConfiguration"`
}

