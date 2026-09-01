package awsbedrockagents


// Experimental.
type AwsBedrockagentDataSource_DataSourceConfigurationWebConfigurationCrawlerConfigurationProperty struct {
	// crawler_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#crawler_limits AwsBedrockagentDataSource#crawler_limits}
	// Experimental.
	CrawlerLimits interface{} `field:"optional" json:"crawlerLimits" yaml:"crawlerLimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#exclusion_filters AwsBedrockagentDataSource#exclusion_filters}.
	// Experimental.
	ExclusionFilters *[]*string `field:"optional" json:"exclusionFilters" yaml:"exclusionFilters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#inclusion_filters AwsBedrockagentDataSource#inclusion_filters}.
	// Experimental.
	InclusionFilters *[]*string `field:"optional" json:"inclusionFilters" yaml:"inclusionFilters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#scope AwsBedrockagentDataSource#scope}.
	// Experimental.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#user_agent AwsBedrockagentDataSource#user_agent}.
	// Experimental.
	UserAgent *string `field:"optional" json:"userAgent" yaml:"userAgent"`
}

