package awsbedrockagents


// Experimental.
type TfDataSource_DataSourceConfigurationSalesforceConfigurationCrawlerConfigurationFilterConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#type TfDataSource#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// pattern_object_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#pattern_object_filter TfDataSource#pattern_object_filter}
	// Experimental.
	PatternObjectFilter interface{} `field:"optional" json:"patternObjectFilter" yaml:"patternObjectFilter"`
}

