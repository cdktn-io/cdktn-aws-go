package awsbedrockagents


// Experimental.
type TfDataSource_SharePointConfigurationProperty struct {
	// crawler_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#crawler_configuration TfDataSource#crawler_configuration}
	// Experimental.
	CrawlerConfiguration interface{} `field:"optional" json:"crawlerConfiguration" yaml:"crawlerConfiguration"`
	// source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#source_configuration TfDataSource#source_configuration}
	// Experimental.
	SourceConfiguration interface{} `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

