package awsbedrockagents


// Experimental.
type AwsBedrockagentDataSource_DataSourceConfigurationSharePointConfigurationCrawlerConfigurationFilterConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#type AwsBedrockagentDataSource#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// pattern_object_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#pattern_object_filter AwsBedrockagentDataSource#pattern_object_filter}
	// Experimental.
	PatternObjectFilter interface{} `field:"optional" json:"patternObjectFilter" yaml:"patternObjectFilter"`
}

