package awsbedrockagents


// Experimental.
type AwsBedrockagentDataSource_HierarchicalChunkingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#overlap_tokens AwsBedrockagentDataSource#overlap_tokens}.
	// Experimental.
	OverlapTokens *float64 `field:"required" json:"overlapTokens" yaml:"overlapTokens"`
	// level_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#level_configuration AwsBedrockagentDataSource#level_configuration}
	// Experimental.
	LevelConfiguration interface{} `field:"optional" json:"levelConfiguration" yaml:"levelConfiguration"`
}

